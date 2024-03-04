import "htmx.org"
import * as bootstrap from "bootstrap"
import "bootstrap/dist/css/bootstrap.min.css"
import "./color-modes"
import "./navbar-fixed.css"
import "hyperscript.org"

window.bootstrap = bootstrap

window.videoPage = function () {

    let devices
    let cameraSelect = document.getElementById("selectCamera")
    const localVideo = document.getElementById('localVideo')
    const remoteVideo = document.getElementById('remoteVideo')

    let recordButton = document.getElementById("recordButton")
    let resolutionSelect = document.getElementById("selectResolution")
    const uuid = document.getElementById("uuid").value

    recordButton.onclick = captureCamera

    let hasMicrophone = false;

    document.getElementById('startButton').addEventListener("click", async function () {
        await listDevices()
    })



    let serverConnection;
    let serverConnectionPromis
    /**
     * @returns Promise<WebSocket>
     */
    async function getWs() {

        if (serverConnectionPromis) {
            return serverConnectionPromis;
        }


        serverConnectionPromis = new Promise((resolve, reject) => {
            if (serverConnection) {
                resolve(serverConnection)
                return
            }


            const connect = function (resolve) {
                const wsHost = window.location.hostname + (window.location.port ? ":" + window.location.port : "")
                const protocol = window.location.protocol == "http:" ? "ws" : "wss"

                const wsUrl = protocol + "://" + wsHost + "/video/room/ws"

                const localServerConnection = new WebSocket(wsUrl);
                localServerConnection.onmessage = async function (message) {


                    const signal = JSON.parse(message.data)
                    console.log("signal", signal)

                    if (signal.uuid == uuid) {
                        return
                    }


                    if (signal.sdp) {
                        await localPeer.setRemoteDescription(signal.sdp)

                        if (signal.sdp.type == 'offer') {
                            const localDescription = await localPeer.createAnswer()
                            await localPeer.setLocalDescription(localDescription)
                        }
                    }

                    if (signal.ice) {
                        if (localPeer.remoteDescription) {
                            await localPeer.addIceCandidate(signal.ice)
                        }
                    }
                }

                localServerConnection.onclose = async function (e) {
                    serverConnectionPromis = undefined
                    serverConnection = undefined
                    connect()
                }

                localServerConnection.onopen = async function (e) {
                    serverConnection = localServerConnection
                    if (resolve) {
                        resolve(localServerConnection)
                    }
                }
            }

            connect(resolve)


        })

        return serverConnectionPromis
    }

    const peerConnectionConfig = {
        iceServers: [
            { 'urls': 'stun:stun.stunprotocol.org:3478' },
            { 'urls': 'stun:stun.l.google.com:19302' },
        ]
    }

    const localPeer = new RTCPeerConnection(peerConnectionConfig)
    localPeer.onicecandidate = async function (e) {
        console.log("localPeer onicecandidate", e)

        if (e.candidate != null) {
            const ws = await getWs()
            ws.send(JSON.stringify({ ice: e.candidate, uuid: uuid }));
        }
    }

    /**
     * @param {RTCTrackEvent} e 
     */
    localPeer.ontrack = async function (e) {
        console.log("localPeer ontrack", e)
        remoteVideo.srcObject = e.streams[0];
    }

    localPeer.onnegotiationneeded = async function (e) {
        console.log("localPeer onnegotiationneeded", e)
    }

    async function askForPermissions() {
        let stream;
        try {
            const constraints = { video: true, audio: true }
            stream = await navigator.mediaDevices.getUserMedia(constraints)
        } catch (error) {
            console.log(error)
            const constraints = { video: true, audio: false }
            stream = await navigator.mediaDevices.getUserMedia(constraints)
            hasMicrophone = false
        }
        closeStream(stream)
    }

    function closeStream(stream) {
        try {
            if (stream) {
                stream.getTracks().forEach(track => track.stop());
            }
        } catch (e) {
            alert(e.message);
        }
    }

    async function listDevices() {
        const devices = await getCameraDevices()
        for (let index = 0; index < devices.length; index++) {
            const device = devices[index];

            const option = document.createElement("option");
            option.value = device.deviceId
            option.text = device.label != "" ? device.label : "Camera " + index
            cameraSelect.add(option);
        }
    }

    async function getCameraDevices() {
        await askForPermissions()
        devices = await navigator.mediaDevices.enumerateDevices()
        const cameraDevices = []
        for (let i = 0; i < devices.length; i++) {
            const device = devices[i]
            if (device.kind == 'videoinput') {
                cameraDevices.push(device)
            }
        }
        return cameraDevices;
    }

    async function captureCamera() {
        let constraints = {
            audio: false,
            video: true
        }

        if (cameraSelect.selectedOptions[0]) {
            var device = devices[cameraSelect.selectedIndex];
            var deviceID = device.deviceId;
            ;
            constraints = {
                video: { deviceId: deviceID },
                audio: hasMicrophone
            }
        }
        if (resolutionSelect.selectedOptions[0]) {
            let desiredWidth = resolutionSelect.selectedOptions[0].label.split("x")[0];
            let desiredHeight = resolutionSelect.selectedOptions[0].label.split("x")[1];
            if (constraints["video"]["deviceId"]) {
                constraints["video"]["width"] = desiredWidth;
                constraints["video"]["height"] = desiredHeight;
            } else {
                constraints = { width: desiredWidth, height: desiredHeight };
            }
        }

        navigator.mediaDevices.getUserMedia(constraints).then(function (stream) {

            for (const track of stream.getTracks()) {
                localPeer.addTrack(track, stream);
            }

            localVideo.srcObject = stream;


        }).catch(function (error) {
            alert('Unable to capture your camera. Please check console logs.');
            console.error(error);
        });
    }

    async function enterInRoom() {
        if (!localPeer.localDescription) {
            const offerDescription = await localPeer.createOffer()
            await localPeer.setLocalDescription(offerDescription)
        }

        const ws = await getWs()
        console.log('senf to ws')
        ws.send(JSON.stringify({ sdp: localPeer.localDescription, uuid: uuid }))
    }


    document.getElementById('enterInRoom')
        .addEventListener("click", async function (e) {
            return enterInRoom()

        })
}