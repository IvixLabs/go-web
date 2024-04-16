import "./shared"
import {createPeer} from "./video/peer";


function main() {
    const uuid = document.getElementById("uuid").value

    const startButton = document.getElementById('startButton')
    const recordButton = document.getElementById("recordButton")
    const stopButton = document.getElementById("stopButton")

    let devices

    const localVideo = document.getElementById('localVideo')
    const remoteVideo = document.getElementById('remoteVideo')
    const resolutionSelect = document.getElementById("selectResolution")
    const cameraSelect = document.getElementById("selectCamera")
    const msgInput = document.getElementById("msgInput")
    const msgButton = document.getElementById("msgButton")
    const msgChat = document.getElementById("msgChat")

    const localPeer = createPeer(remoteVideo, uuid)

    const dataChannel = localPeer.createDataChannel("messages",
        {negotiated: false}
    )

    dataChannel.onmessage = function (event) {
        console.log(event)

        let newParagraph = document.createElement("p");
        let textNode = document.createTextNode(event.data);
        newParagraph.appendChild(textNode);
        msgChat.appendChild(newParagraph)
    }

    msgButton.onclick = function () {
        const msg = msgInput.value
        console.log('send msg')
        console.log(msg)
        dataChannel.send(msg)
    }

    recordButton.onclick = async function () {
        return captureCamera(localPeer)
    }

    stopButton.onclick = async function () {
        return stopCamera(localPeer)
    }

    let hasMicrophone = true;

    startButton.onclick = async function () {
        startButton.disabled = true
        await listDevices()
        recordButton.classList.remove("d-none")
        stopButton.classList.remove("d-none")
        startButton.classList.add("d-none")

    }


    async function askForPermissions() {
        let stream;
        try {
            const constraints = {video: true, audio: true}
            stream = await navigator.mediaDevices.getUserMedia(constraints)
        } catch (error) {
            console.log(error)
            const constraints = {video: true, audio: false}
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
            option.text = device.label !== "" ? device.label : "Camera " + index
            cameraSelect.add(option);
        }
    }

    async function getCameraDevices() {
        await askForPermissions()
        devices = await navigator.mediaDevices.enumerateDevices()
        const cameraDevices = []
        for (let i = 0; i < devices.length; i++) {
            const device = devices[i]
            if (device.kind === 'videoinput') {
                cameraDevices.push(device)
            }
        }
        return cameraDevices;
    }


    /**
     * @param peer {RTCPeerConnection}
     * @returns {Promise<void>}
     */
    async function stopCamera(peer) {
        peer.getSenders().forEach(function (sender) {
            sender.track.stop()
        })
    }

    /**
     * @param peer {RTCPeerConnection}
     * @returns {Promise<void>}
     */
    async function captureCamera(peer) {
        let constraints = {
            audio: false,
            video: true
        }

        if (cameraSelect.selectedOptions[0]) {
            let device = devices[cameraSelect.selectedIndex];
            let deviceID = device.deviceId;

            constraints = {
                video: {deviceId: deviceID},
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
                constraints = {width: desiredWidth, height: desiredHeight};
            }
        }

        navigator.mediaDevices.getUserMedia(constraints).then(function (stream) {

            stopCamera(peer)

            console.log(peer.getSenders())
            peer.getSenders()

            let foundSender = false
            peer.getSenders().forEach(function (sender) {
                sender.setStreams(stream)
                //     if (sender.track.kind === track.kind) {
                //         sender.replaceTrack(track)
                foundSender = true
                //     }
            })

            if (!foundSender) {
                for (const track of stream.getTracks()) {
                    peer.addTrack(track, stream)
                }
            } else {
                for (const track of stream.getTracks()) {
                    peer.getSenders().forEach(function (sender) {
                        if (sender.track.kind === track.kind) {
                            sender.replaceTrack(track)
                        }
                    })
                }
            }

            // for (const track of stream.getTracks()) {

            // let foundSender = false
            // peer.getSenders().forEach(function (sender) {
            //     if (sender.track.kind === track.kind) {
            //         sender.replaceTrack(track)
            //         foundSender = true
            //     }
            // })

            // if (!foundSender) {
            // peer.addTrack(track, stream)
            // }
            // }

            localVideo.srcObject = stream;


        }).catch(function (error) {
            alert('Unable to capture your camera. Please check console logs.');
            console.error(error);
        });
    }

}

(function () {
    main()
})();