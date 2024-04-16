"use strict";
(self["webpackChunk"] = self["webpackChunk"] || []).push([["video"],{

/***/ "./video.js":
/*!******************!*\
  !*** ./video.js ***!
  \******************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _shared__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./shared */ "./shared.js");
/* harmony import */ var _video_peer__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./video/peer */ "./video/peer.js");




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

    const localPeer = (0,_video_peer__WEBPACK_IMPORTED_MODULE_1__.createPeer)(remoteVideo, uuid)

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

/***/ }),

/***/ "./video/peer.js":
/*!***********************!*\
  !*** ./video/peer.js ***!
  \***********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   createPeer: () => (/* binding */ createPeer)
/* harmony export */ });
/* harmony import */ var _ws__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ws */ "./video/ws.js");


/**
 * @param remoteVideo {HTMLMediaElement}
 * @param uuid {string}
 * @returns {RTCPeerConnection}
 */
function createPeer(remoteVideo, uuid) {
    const peerConnectionConfig = {
        iceServers: [
            {'urls': 'stun:stun.stunprotocol.org:3478'},
            {'urls': 'stun:stun.l.google.com:19302'},
        ]
    }

    const peer = new RTCPeerConnection(peerConnectionConfig)

    peer.onicecandidate = async function (e) {
        console.log("localPeer onicecandidate", e)

        if (e.candidate != null) {
            const ws = await (0,_ws__WEBPACK_IMPORTED_MODULE_0__.getWs)(getOnMessageWsCallback(peer, uuid))
            ws.send(JSON.stringify({ice: e.candidate, uuid: uuid}));
        }
    }

    peer.onnegotiationneeded = async function (e) {
        console.log("localPeer onnegotiationneeded", e)
        return enterInRoom(this, uuid)
    }

    /**
     * @param {RTCTrackEvent} e
     */
    peer.ontrack = async function (e) {
        console.log("localPeer ontrack", e)
        remoteVideo.srcObject = e.streams[0];
    }

    return peer
}

/**
 * @param peer {RTCPeerConnection}
 * @param uuid {string}
 * @returns {Promise<void>}
 */
async function enterInRoom(peer, uuid) {
    if (!peer.localDescription) {
        const offerDescription = await peer.createOffer()
        await peer.setLocalDescription(offerDescription)
    }

    const ws = await (0,_ws__WEBPACK_IMPORTED_MODULE_0__.getWs)(getOnMessageWsCallback(peer, uuid))
    ws.send(JSON.stringify({sdp: peer.localDescription, uuid: uuid}))
}


/**
 * @param peer {RTCPeerConnection}
 * @param uuid {string}
 * @returns {(function(message : MessageEvent): Promise<void>)}
 */
function getOnMessageWsCallback(peer, uuid) {
    return async function (message) {

        const signal = JSON.parse(message.data)
        console.log("signal", signal)

        if (signal.uuid === uuid) {
            return
        }


        if (signal.sdp) {
            if (peer.signalingState !== "stable") {
                await peer.setRemoteDescription(signal.sdp)

                if (signal.sdp.type === 'offer') {
                    const localDescription = await peer.createAnswer()
                    await peer.setLocalDescription(localDescription)
                    await enterInRoom(peer, uuid)
                }
            }
        }

        if (signal.ice) {
            if (peer.remoteDescription) {
                await peer.addIceCandidate(signal.ice)
            }
        }
    }
}

/***/ }),

/***/ "./video/ws.js":
/*!*********************!*\
  !*** ./video/ws.js ***!
  \*********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   getWs: () => (/* binding */ getWs)
/* harmony export */ });
let wsConnection;
let wsConnectionPromise

/**
 *
 * @param onMessageCallback {(function(ev: MessageEvent) : any)}
 * @returns {Promise<WebSocket>}
 */
async function getWs(onMessageCallback) {

    if (wsConnectionPromise) {
        return wsConnectionPromise;
    }


    wsConnectionPromise = new Promise((resolve, reject) => {
        if (wsConnection) {
            resolve(wsConnection)
            return
        }


        const connect = function (resolve) {
            const wsHost = window.location.hostname + (window.location.port ? ":" + window.location.port : "")
            const protocol = window.location.protocol == "http:" ? "ws" : "wss"

            const wsUrl = protocol + "://" + wsHost + "/video/room/ws"

            const localWsConnection = new WebSocket(wsUrl);
            localWsConnection.onmessage = onMessageCallback


            localWsConnection.onclose = async function (e) {
                wsConnectionPromise = undefined
                wsConnection = undefined
                setTimeout(connect, 1000)
            }

            localWsConnection.onopen = async function (e) {
                wsConnection = localWsConnection
                if (resolve) {
                    resolve(localWsConnection)
                }
            }
        }

        connect(resolve)
    })

    return wsConnectionPromise
}

/***/ })

},
/******/ __webpack_require__ => { // webpackRuntimeModules
/******/ var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
/******/ var __webpack_exports__ = (__webpack_exec__("./video.js"));
/******/ }
]);
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoidmlkZW8uYnVuZGxlLmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7OztBQUFpQjtBQUN1Qjs7O0FBR3hDO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBLHNCQUFzQix1REFBVTs7QUFFaEM7QUFDQSxTQUFTO0FBQ1Q7O0FBRUE7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUFFQTs7O0FBR0E7QUFDQTtBQUNBO0FBQ0EsaUNBQWlDO0FBQ2pDO0FBQ0EsVUFBVTtBQUNWO0FBQ0EsaUNBQWlDO0FBQ2pDO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLFVBQVU7QUFDVjtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBLDRCQUE0Qix3QkFBd0I7QUFDcEQ7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0Esd0JBQXdCLG9CQUFvQjtBQUM1QztBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7O0FBR0E7QUFDQSxvQkFBb0I7QUFDcEIsaUJBQWlCO0FBQ2pCO0FBQ0E7QUFDQTtBQUNBO0FBQ0EsU0FBUztBQUNUOztBQUVBO0FBQ0Esb0JBQW9CO0FBQ3BCLGlCQUFpQjtBQUNqQjtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBO0FBQ0Esd0JBQXdCLG1CQUFtQjtBQUMzQztBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxjQUFjO0FBQ2QsK0JBQStCO0FBQy9CO0FBQ0E7O0FBRUE7O0FBRUE7O0FBRUE7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLGFBQWE7O0FBRWI7QUFDQTtBQUNBO0FBQ0E7QUFDQSxjQUFjO0FBQ2Q7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLHFCQUFxQjtBQUNyQjtBQUNBOztBQUVBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLGdCQUFnQjs7QUFFaEI7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7OztBQUdBLFNBQVM7QUFDVDtBQUNBO0FBQ0EsU0FBUztBQUNUOztBQUVBOztBQUVBO0FBQ0E7QUFDQSxDQUFDOzs7Ozs7Ozs7Ozs7Ozs7QUNuTjBCOztBQUUzQjtBQUNBLHVCQUF1QjtBQUN2QixnQkFBZ0I7QUFDaEIsYUFBYTtBQUNiO0FBQ087QUFDUDtBQUNBO0FBQ0EsYUFBYSwwQ0FBMEM7QUFDdkQsYUFBYSx1Q0FBdUM7QUFDcEQ7QUFDQTs7QUFFQTs7QUFFQTtBQUNBOztBQUVBO0FBQ0EsNkJBQTZCLDBDQUFLO0FBQ2xDLG9DQUFvQyw2QkFBNkI7QUFDakU7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7QUFDQTs7QUFFQTtBQUNBLGVBQWUsZUFBZTtBQUM5QjtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7O0FBRUE7QUFDQSxnQkFBZ0I7QUFDaEIsZ0JBQWdCO0FBQ2hCLGFBQWE7QUFDYjtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUEscUJBQXFCLDBDQUFLO0FBQzFCLDRCQUE0Qix1Q0FBdUM7QUFDbkU7OztBQUdBO0FBQ0EsZ0JBQWdCO0FBQ2hCLGdCQUFnQjtBQUNoQixhQUFhO0FBQ2I7QUFDQTtBQUNBOztBQUVBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOzs7QUFHQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7Ozs7Ozs7Ozs7Ozs7O0FDNUZBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBLDZCQUE2QjtBQUM3QixhQUFhO0FBQ2I7QUFDTzs7QUFFUDtBQUNBO0FBQ0E7OztBQUdBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7OztBQUdBO0FBQ0E7QUFDQTs7QUFFQTs7QUFFQTtBQUNBOzs7QUFHQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0EsS0FBSzs7QUFFTDtBQUNBIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vLy4vdmlkZW8uanMiLCJ3ZWJwYWNrOi8vLy4vdmlkZW8vcGVlci5qcyIsIndlYnBhY2s6Ly8vLi92aWRlby93cy5qcyJdLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgXCIuL3NoYXJlZFwiXG5pbXBvcnQge2NyZWF0ZVBlZXJ9IGZyb20gXCIuL3ZpZGVvL3BlZXJcIjtcblxuXG5mdW5jdGlvbiBtYWluKCkge1xuICAgIGNvbnN0IHV1aWQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInV1aWRcIikudmFsdWVcblxuICAgIGNvbnN0IHN0YXJ0QnV0dG9uID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3N0YXJ0QnV0dG9uJylcbiAgICBjb25zdCByZWNvcmRCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInJlY29yZEJ1dHRvblwiKVxuICAgIGNvbnN0IHN0b3BCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcInN0b3BCdXR0b25cIilcblxuICAgIGxldCBkZXZpY2VzXG5cbiAgICBjb25zdCBsb2NhbFZpZGVvID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ2xvY2FsVmlkZW8nKVxuICAgIGNvbnN0IHJlbW90ZVZpZGVvID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoJ3JlbW90ZVZpZGVvJylcbiAgICBjb25zdCByZXNvbHV0aW9uU2VsZWN0ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJzZWxlY3RSZXNvbHV0aW9uXCIpXG4gICAgY29uc3QgY2FtZXJhU2VsZWN0ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQoXCJzZWxlY3RDYW1lcmFcIilcbiAgICBjb25zdCBtc2dJbnB1dCA9IGRvY3VtZW50LmdldEVsZW1lbnRCeUlkKFwibXNnSW5wdXRcIilcbiAgICBjb25zdCBtc2dCdXR0b24gPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcIm1zZ0J1dHRvblwiKVxuICAgIGNvbnN0IG1zZ0NoYXQgPSBkb2N1bWVudC5nZXRFbGVtZW50QnlJZChcIm1zZ0NoYXRcIilcblxuICAgIGNvbnN0IGxvY2FsUGVlciA9IGNyZWF0ZVBlZXIocmVtb3RlVmlkZW8sIHV1aWQpXG5cbiAgICBjb25zdCBkYXRhQ2hhbm5lbCA9IGxvY2FsUGVlci5jcmVhdGVEYXRhQ2hhbm5lbChcIm1lc3NhZ2VzXCIsXG4gICAgICAgIHtuZWdvdGlhdGVkOiBmYWxzZX1cbiAgICApXG5cbiAgICBkYXRhQ2hhbm5lbC5vbm1lc3NhZ2UgPSBmdW5jdGlvbiAoZXZlbnQpIHtcbiAgICAgICAgY29uc29sZS5sb2coZXZlbnQpXG5cbiAgICAgICAgbGV0IG5ld1BhcmFncmFwaCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJwXCIpO1xuICAgICAgICBsZXQgdGV4dE5vZGUgPSBkb2N1bWVudC5jcmVhdGVUZXh0Tm9kZShldmVudC5kYXRhKTtcbiAgICAgICAgbmV3UGFyYWdyYXBoLmFwcGVuZENoaWxkKHRleHROb2RlKTtcbiAgICAgICAgbXNnQ2hhdC5hcHBlbmRDaGlsZChuZXdQYXJhZ3JhcGgpXG4gICAgfVxuXG4gICAgbXNnQnV0dG9uLm9uY2xpY2sgPSBmdW5jdGlvbiAoKSB7XG4gICAgICAgIGNvbnN0IG1zZyA9IG1zZ0lucHV0LnZhbHVlXG4gICAgICAgIGNvbnNvbGUubG9nKCdzZW5kIG1zZycpXG4gICAgICAgIGNvbnNvbGUubG9nKG1zZylcbiAgICAgICAgZGF0YUNoYW5uZWwuc2VuZChtc2cpXG4gICAgfVxuXG4gICAgcmVjb3JkQnV0dG9uLm9uY2xpY2sgPSBhc3luYyBmdW5jdGlvbiAoKSB7XG4gICAgICAgIHJldHVybiBjYXB0dXJlQ2FtZXJhKGxvY2FsUGVlcilcbiAgICB9XG5cbiAgICBzdG9wQnV0dG9uLm9uY2xpY2sgPSBhc3luYyBmdW5jdGlvbiAoKSB7XG4gICAgICAgIHJldHVybiBzdG9wQ2FtZXJhKGxvY2FsUGVlcilcbiAgICB9XG5cbiAgICBsZXQgaGFzTWljcm9waG9uZSA9IHRydWU7XG5cbiAgICBzdGFydEJ1dHRvbi5vbmNsaWNrID0gYXN5bmMgZnVuY3Rpb24gKCkge1xuICAgICAgICBzdGFydEJ1dHRvbi5kaXNhYmxlZCA9IHRydWVcbiAgICAgICAgYXdhaXQgbGlzdERldmljZXMoKVxuICAgICAgICByZWNvcmRCdXR0b24uY2xhc3NMaXN0LnJlbW92ZShcImQtbm9uZVwiKVxuICAgICAgICBzdG9wQnV0dG9uLmNsYXNzTGlzdC5yZW1vdmUoXCJkLW5vbmVcIilcbiAgICAgICAgc3RhcnRCdXR0b24uY2xhc3NMaXN0LmFkZChcImQtbm9uZVwiKVxuXG4gICAgfVxuXG5cbiAgICBhc3luYyBmdW5jdGlvbiBhc2tGb3JQZXJtaXNzaW9ucygpIHtcbiAgICAgICAgbGV0IHN0cmVhbTtcbiAgICAgICAgdHJ5IHtcbiAgICAgICAgICAgIGNvbnN0IGNvbnN0cmFpbnRzID0ge3ZpZGVvOiB0cnVlLCBhdWRpbzogdHJ1ZX1cbiAgICAgICAgICAgIHN0cmVhbSA9IGF3YWl0IG5hdmlnYXRvci5tZWRpYURldmljZXMuZ2V0VXNlck1lZGlhKGNvbnN0cmFpbnRzKVxuICAgICAgICB9IGNhdGNoIChlcnJvcikge1xuICAgICAgICAgICAgY29uc29sZS5sb2coZXJyb3IpXG4gICAgICAgICAgICBjb25zdCBjb25zdHJhaW50cyA9IHt2aWRlbzogdHJ1ZSwgYXVkaW86IGZhbHNlfVxuICAgICAgICAgICAgc3RyZWFtID0gYXdhaXQgbmF2aWdhdG9yLm1lZGlhRGV2aWNlcy5nZXRVc2VyTWVkaWEoY29uc3RyYWludHMpXG4gICAgICAgICAgICBoYXNNaWNyb3Bob25lID0gZmFsc2VcbiAgICAgICAgfVxuICAgICAgICBjbG9zZVN0cmVhbShzdHJlYW0pXG4gICAgfVxuXG4gICAgZnVuY3Rpb24gY2xvc2VTdHJlYW0oc3RyZWFtKSB7XG4gICAgICAgIHRyeSB7XG4gICAgICAgICAgICBpZiAoc3RyZWFtKSB7XG4gICAgICAgICAgICAgICAgc3RyZWFtLmdldFRyYWNrcygpLmZvckVhY2godHJhY2sgPT4gdHJhY2suc3RvcCgpKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgfSBjYXRjaCAoZSkge1xuICAgICAgICAgICAgYWxlcnQoZS5tZXNzYWdlKTtcbiAgICAgICAgfVxuICAgIH1cblxuICAgIGFzeW5jIGZ1bmN0aW9uIGxpc3REZXZpY2VzKCkge1xuICAgICAgICBjb25zdCBkZXZpY2VzID0gYXdhaXQgZ2V0Q2FtZXJhRGV2aWNlcygpXG4gICAgICAgIGZvciAobGV0IGluZGV4ID0gMDsgaW5kZXggPCBkZXZpY2VzLmxlbmd0aDsgaW5kZXgrKykge1xuICAgICAgICAgICAgY29uc3QgZGV2aWNlID0gZGV2aWNlc1tpbmRleF07XG5cbiAgICAgICAgICAgIGNvbnN0IG9wdGlvbiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJvcHRpb25cIik7XG4gICAgICAgICAgICBvcHRpb24udmFsdWUgPSBkZXZpY2UuZGV2aWNlSWRcbiAgICAgICAgICAgIG9wdGlvbi50ZXh0ID0gZGV2aWNlLmxhYmVsICE9PSBcIlwiID8gZGV2aWNlLmxhYmVsIDogXCJDYW1lcmEgXCIgKyBpbmRleFxuICAgICAgICAgICAgY2FtZXJhU2VsZWN0LmFkZChvcHRpb24pO1xuICAgICAgICB9XG4gICAgfVxuXG4gICAgYXN5bmMgZnVuY3Rpb24gZ2V0Q2FtZXJhRGV2aWNlcygpIHtcbiAgICAgICAgYXdhaXQgYXNrRm9yUGVybWlzc2lvbnMoKVxuICAgICAgICBkZXZpY2VzID0gYXdhaXQgbmF2aWdhdG9yLm1lZGlhRGV2aWNlcy5lbnVtZXJhdGVEZXZpY2VzKClcbiAgICAgICAgY29uc3QgY2FtZXJhRGV2aWNlcyA9IFtdXG4gICAgICAgIGZvciAobGV0IGkgPSAwOyBpIDwgZGV2aWNlcy5sZW5ndGg7IGkrKykge1xuICAgICAgICAgICAgY29uc3QgZGV2aWNlID0gZGV2aWNlc1tpXVxuICAgICAgICAgICAgaWYgKGRldmljZS5raW5kID09PSAndmlkZW9pbnB1dCcpIHtcbiAgICAgICAgICAgICAgICBjYW1lcmFEZXZpY2VzLnB1c2goZGV2aWNlKVxuICAgICAgICAgICAgfVxuICAgICAgICB9XG4gICAgICAgIHJldHVybiBjYW1lcmFEZXZpY2VzO1xuICAgIH1cblxuXG4gICAgLyoqXG4gICAgICogQHBhcmFtIHBlZXIge1JUQ1BlZXJDb25uZWN0aW9ufVxuICAgICAqIEByZXR1cm5zIHtQcm9taXNlPHZvaWQ+fVxuICAgICAqL1xuICAgIGFzeW5jIGZ1bmN0aW9uIHN0b3BDYW1lcmEocGVlcikge1xuICAgICAgICBwZWVyLmdldFNlbmRlcnMoKS5mb3JFYWNoKGZ1bmN0aW9uIChzZW5kZXIpIHtcbiAgICAgICAgICAgIHNlbmRlci50cmFjay5zdG9wKClcbiAgICAgICAgfSlcbiAgICB9XG5cbiAgICAvKipcbiAgICAgKiBAcGFyYW0gcGVlciB7UlRDUGVlckNvbm5lY3Rpb259XG4gICAgICogQHJldHVybnMge1Byb21pc2U8dm9pZD59XG4gICAgICovXG4gICAgYXN5bmMgZnVuY3Rpb24gY2FwdHVyZUNhbWVyYShwZWVyKSB7XG4gICAgICAgIGxldCBjb25zdHJhaW50cyA9IHtcbiAgICAgICAgICAgIGF1ZGlvOiBmYWxzZSxcbiAgICAgICAgICAgIHZpZGVvOiB0cnVlXG4gICAgICAgIH1cblxuICAgICAgICBpZiAoY2FtZXJhU2VsZWN0LnNlbGVjdGVkT3B0aW9uc1swXSkge1xuICAgICAgICAgICAgbGV0IGRldmljZSA9IGRldmljZXNbY2FtZXJhU2VsZWN0LnNlbGVjdGVkSW5kZXhdO1xuICAgICAgICAgICAgbGV0IGRldmljZUlEID0gZGV2aWNlLmRldmljZUlkO1xuXG4gICAgICAgICAgICBjb25zdHJhaW50cyA9IHtcbiAgICAgICAgICAgICAgICB2aWRlbzoge2RldmljZUlkOiBkZXZpY2VJRH0sXG4gICAgICAgICAgICAgICAgYXVkaW86IGhhc01pY3JvcGhvbmVcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgICAgICBpZiAocmVzb2x1dGlvblNlbGVjdC5zZWxlY3RlZE9wdGlvbnNbMF0pIHtcbiAgICAgICAgICAgIGxldCBkZXNpcmVkV2lkdGggPSByZXNvbHV0aW9uU2VsZWN0LnNlbGVjdGVkT3B0aW9uc1swXS5sYWJlbC5zcGxpdChcInhcIilbMF07XG4gICAgICAgICAgICBsZXQgZGVzaXJlZEhlaWdodCA9IHJlc29sdXRpb25TZWxlY3Quc2VsZWN0ZWRPcHRpb25zWzBdLmxhYmVsLnNwbGl0KFwieFwiKVsxXTtcbiAgICAgICAgICAgIGlmIChjb25zdHJhaW50c1tcInZpZGVvXCJdW1wiZGV2aWNlSWRcIl0pIHtcbiAgICAgICAgICAgICAgICBjb25zdHJhaW50c1tcInZpZGVvXCJdW1wid2lkdGhcIl0gPSBkZXNpcmVkV2lkdGg7XG4gICAgICAgICAgICAgICAgY29uc3RyYWludHNbXCJ2aWRlb1wiXVtcImhlaWdodFwiXSA9IGRlc2lyZWRIZWlnaHQ7XG4gICAgICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgICAgICAgIGNvbnN0cmFpbnRzID0ge3dpZHRoOiBkZXNpcmVkV2lkdGgsIGhlaWdodDogZGVzaXJlZEhlaWdodH07XG4gICAgICAgICAgICB9XG4gICAgICAgIH1cblxuICAgICAgICBuYXZpZ2F0b3IubWVkaWFEZXZpY2VzLmdldFVzZXJNZWRpYShjb25zdHJhaW50cykudGhlbihmdW5jdGlvbiAoc3RyZWFtKSB7XG5cbiAgICAgICAgICAgIHN0b3BDYW1lcmEocGVlcilcblxuICAgICAgICAgICAgY29uc29sZS5sb2cocGVlci5nZXRTZW5kZXJzKCkpXG4gICAgICAgICAgICBwZWVyLmdldFNlbmRlcnMoKVxuXG4gICAgICAgICAgICBsZXQgZm91bmRTZW5kZXIgPSBmYWxzZVxuICAgICAgICAgICAgcGVlci5nZXRTZW5kZXJzKCkuZm9yRWFjaChmdW5jdGlvbiAoc2VuZGVyKSB7XG4gICAgICAgICAgICAgICAgc2VuZGVyLnNldFN0cmVhbXMoc3RyZWFtKVxuICAgICAgICAgICAgICAgIC8vICAgICBpZiAoc2VuZGVyLnRyYWNrLmtpbmQgPT09IHRyYWNrLmtpbmQpIHtcbiAgICAgICAgICAgICAgICAvLyAgICAgICAgIHNlbmRlci5yZXBsYWNlVHJhY2sodHJhY2spXG4gICAgICAgICAgICAgICAgZm91bmRTZW5kZXIgPSB0cnVlXG4gICAgICAgICAgICAgICAgLy8gICAgIH1cbiAgICAgICAgICAgIH0pXG5cbiAgICAgICAgICAgIGlmICghZm91bmRTZW5kZXIpIHtcbiAgICAgICAgICAgICAgICBmb3IgKGNvbnN0IHRyYWNrIG9mIHN0cmVhbS5nZXRUcmFja3MoKSkge1xuICAgICAgICAgICAgICAgICAgICBwZWVyLmFkZFRyYWNrKHRyYWNrLCBzdHJlYW0pXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgICAgICBmb3IgKGNvbnN0IHRyYWNrIG9mIHN0cmVhbS5nZXRUcmFja3MoKSkge1xuICAgICAgICAgICAgICAgICAgICBwZWVyLmdldFNlbmRlcnMoKS5mb3JFYWNoKGZ1bmN0aW9uIChzZW5kZXIpIHtcbiAgICAgICAgICAgICAgICAgICAgICAgIGlmIChzZW5kZXIudHJhY2sua2luZCA9PT0gdHJhY2sua2luZCkge1xuICAgICAgICAgICAgICAgICAgICAgICAgICAgIHNlbmRlci5yZXBsYWNlVHJhY2sodHJhY2spXG4gICAgICAgICAgICAgICAgICAgICAgICB9XG4gICAgICAgICAgICAgICAgICAgIH0pXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfVxuXG4gICAgICAgICAgICAvLyBmb3IgKGNvbnN0IHRyYWNrIG9mIHN0cmVhbS5nZXRUcmFja3MoKSkge1xuXG4gICAgICAgICAgICAvLyBsZXQgZm91bmRTZW5kZXIgPSBmYWxzZVxuICAgICAgICAgICAgLy8gcGVlci5nZXRTZW5kZXJzKCkuZm9yRWFjaChmdW5jdGlvbiAoc2VuZGVyKSB7XG4gICAgICAgICAgICAvLyAgICAgaWYgKHNlbmRlci50cmFjay5raW5kID09PSB0cmFjay5raW5kKSB7XG4gICAgICAgICAgICAvLyAgICAgICAgIHNlbmRlci5yZXBsYWNlVHJhY2sodHJhY2spXG4gICAgICAgICAgICAvLyAgICAgICAgIGZvdW5kU2VuZGVyID0gdHJ1ZVxuICAgICAgICAgICAgLy8gICAgIH1cbiAgICAgICAgICAgIC8vIH0pXG5cbiAgICAgICAgICAgIC8vIGlmICghZm91bmRTZW5kZXIpIHtcbiAgICAgICAgICAgIC8vIHBlZXIuYWRkVHJhY2sodHJhY2ssIHN0cmVhbSlcbiAgICAgICAgICAgIC8vIH1cbiAgICAgICAgICAgIC8vIH1cblxuICAgICAgICAgICAgbG9jYWxWaWRlby5zcmNPYmplY3QgPSBzdHJlYW07XG5cblxuICAgICAgICB9KS5jYXRjaChmdW5jdGlvbiAoZXJyb3IpIHtcbiAgICAgICAgICAgIGFsZXJ0KCdVbmFibGUgdG8gY2FwdHVyZSB5b3VyIGNhbWVyYS4gUGxlYXNlIGNoZWNrIGNvbnNvbGUgbG9ncy4nKTtcbiAgICAgICAgICAgIGNvbnNvbGUuZXJyb3IoZXJyb3IpO1xuICAgICAgICB9KTtcbiAgICB9XG5cbn1cblxuKGZ1bmN0aW9uICgpIHtcbiAgICBtYWluKClcbn0pKCk7IiwiaW1wb3J0IHtnZXRXc30gZnJvbSBcIi4vd3NcIjtcblxuLyoqXG4gKiBAcGFyYW0gcmVtb3RlVmlkZW8ge0hUTUxNZWRpYUVsZW1lbnR9XG4gKiBAcGFyYW0gdXVpZCB7c3RyaW5nfVxuICogQHJldHVybnMge1JUQ1BlZXJDb25uZWN0aW9ufVxuICovXG5leHBvcnQgZnVuY3Rpb24gY3JlYXRlUGVlcihyZW1vdGVWaWRlbywgdXVpZCkge1xuICAgIGNvbnN0IHBlZXJDb25uZWN0aW9uQ29uZmlnID0ge1xuICAgICAgICBpY2VTZXJ2ZXJzOiBbXG4gICAgICAgICAgICB7J3VybHMnOiAnc3R1bjpzdHVuLnN0dW5wcm90b2NvbC5vcmc6MzQ3OCd9LFxuICAgICAgICAgICAgeyd1cmxzJzogJ3N0dW46c3R1bi5sLmdvb2dsZS5jb206MTkzMDInfSxcbiAgICAgICAgXVxuICAgIH1cblxuICAgIGNvbnN0IHBlZXIgPSBuZXcgUlRDUGVlckNvbm5lY3Rpb24ocGVlckNvbm5lY3Rpb25Db25maWcpXG5cbiAgICBwZWVyLm9uaWNlY2FuZGlkYXRlID0gYXN5bmMgZnVuY3Rpb24gKGUpIHtcbiAgICAgICAgY29uc29sZS5sb2coXCJsb2NhbFBlZXIgb25pY2VjYW5kaWRhdGVcIiwgZSlcblxuICAgICAgICBpZiAoZS5jYW5kaWRhdGUgIT0gbnVsbCkge1xuICAgICAgICAgICAgY29uc3Qgd3MgPSBhd2FpdCBnZXRXcyhnZXRPbk1lc3NhZ2VXc0NhbGxiYWNrKHBlZXIsIHV1aWQpKVxuICAgICAgICAgICAgd3Muc2VuZChKU09OLnN0cmluZ2lmeSh7aWNlOiBlLmNhbmRpZGF0ZSwgdXVpZDogdXVpZH0pKTtcbiAgICAgICAgfVxuICAgIH1cblxuICAgIHBlZXIub25uZWdvdGlhdGlvbm5lZWRlZCA9IGFzeW5jIGZ1bmN0aW9uIChlKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKFwibG9jYWxQZWVyIG9ubmVnb3RpYXRpb25uZWVkZWRcIiwgZSlcbiAgICAgICAgcmV0dXJuIGVudGVySW5Sb29tKHRoaXMsIHV1aWQpXG4gICAgfVxuXG4gICAgLyoqXG4gICAgICogQHBhcmFtIHtSVENUcmFja0V2ZW50fSBlXG4gICAgICovXG4gICAgcGVlci5vbnRyYWNrID0gYXN5bmMgZnVuY3Rpb24gKGUpIHtcbiAgICAgICAgY29uc29sZS5sb2coXCJsb2NhbFBlZXIgb250cmFja1wiLCBlKVxuICAgICAgICByZW1vdGVWaWRlby5zcmNPYmplY3QgPSBlLnN0cmVhbXNbMF07XG4gICAgfVxuXG4gICAgcmV0dXJuIHBlZXJcbn1cblxuLyoqXG4gKiBAcGFyYW0gcGVlciB7UlRDUGVlckNvbm5lY3Rpb259XG4gKiBAcGFyYW0gdXVpZCB7c3RyaW5nfVxuICogQHJldHVybnMge1Byb21pc2U8dm9pZD59XG4gKi9cbmFzeW5jIGZ1bmN0aW9uIGVudGVySW5Sb29tKHBlZXIsIHV1aWQpIHtcbiAgICBpZiAoIXBlZXIubG9jYWxEZXNjcmlwdGlvbikge1xuICAgICAgICBjb25zdCBvZmZlckRlc2NyaXB0aW9uID0gYXdhaXQgcGVlci5jcmVhdGVPZmZlcigpXG4gICAgICAgIGF3YWl0IHBlZXIuc2V0TG9jYWxEZXNjcmlwdGlvbihvZmZlckRlc2NyaXB0aW9uKVxuICAgIH1cblxuICAgIGNvbnN0IHdzID0gYXdhaXQgZ2V0V3MoZ2V0T25NZXNzYWdlV3NDYWxsYmFjayhwZWVyLCB1dWlkKSlcbiAgICB3cy5zZW5kKEpTT04uc3RyaW5naWZ5KHtzZHA6IHBlZXIubG9jYWxEZXNjcmlwdGlvbiwgdXVpZDogdXVpZH0pKVxufVxuXG5cbi8qKlxuICogQHBhcmFtIHBlZXIge1JUQ1BlZXJDb25uZWN0aW9ufVxuICogQHBhcmFtIHV1aWQge3N0cmluZ31cbiAqIEByZXR1cm5zIHsoZnVuY3Rpb24obWVzc2FnZSA6IE1lc3NhZ2VFdmVudCk6IFByb21pc2U8dm9pZD4pfVxuICovXG5mdW5jdGlvbiBnZXRPbk1lc3NhZ2VXc0NhbGxiYWNrKHBlZXIsIHV1aWQpIHtcbiAgICByZXR1cm4gYXN5bmMgZnVuY3Rpb24gKG1lc3NhZ2UpIHtcblxuICAgICAgICBjb25zdCBzaWduYWwgPSBKU09OLnBhcnNlKG1lc3NhZ2UuZGF0YSlcbiAgICAgICAgY29uc29sZS5sb2coXCJzaWduYWxcIiwgc2lnbmFsKVxuXG4gICAgICAgIGlmIChzaWduYWwudXVpZCA9PT0gdXVpZCkge1xuICAgICAgICAgICAgcmV0dXJuXG4gICAgICAgIH1cblxuXG4gICAgICAgIGlmIChzaWduYWwuc2RwKSB7XG4gICAgICAgICAgICBpZiAocGVlci5zaWduYWxpbmdTdGF0ZSAhPT0gXCJzdGFibGVcIikge1xuICAgICAgICAgICAgICAgIGF3YWl0IHBlZXIuc2V0UmVtb3RlRGVzY3JpcHRpb24oc2lnbmFsLnNkcClcblxuICAgICAgICAgICAgICAgIGlmIChzaWduYWwuc2RwLnR5cGUgPT09ICdvZmZlcicpIHtcbiAgICAgICAgICAgICAgICAgICAgY29uc3QgbG9jYWxEZXNjcmlwdGlvbiA9IGF3YWl0IHBlZXIuY3JlYXRlQW5zd2VyKClcbiAgICAgICAgICAgICAgICAgICAgYXdhaXQgcGVlci5zZXRMb2NhbERlc2NyaXB0aW9uKGxvY2FsRGVzY3JpcHRpb24pXG4gICAgICAgICAgICAgICAgICAgIGF3YWl0IGVudGVySW5Sb29tKHBlZXIsIHV1aWQpXG4gICAgICAgICAgICAgICAgfVxuICAgICAgICAgICAgfVxuICAgICAgICB9XG5cbiAgICAgICAgaWYgKHNpZ25hbC5pY2UpIHtcbiAgICAgICAgICAgIGlmIChwZWVyLnJlbW90ZURlc2NyaXB0aW9uKSB7XG4gICAgICAgICAgICAgICAgYXdhaXQgcGVlci5hZGRJY2VDYW5kaWRhdGUoc2lnbmFsLmljZSlcbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuICAgIH1cbn0iLCJsZXQgd3NDb25uZWN0aW9uO1xubGV0IHdzQ29ubmVjdGlvblByb21pc2VcblxuLyoqXG4gKlxuICogQHBhcmFtIG9uTWVzc2FnZUNhbGxiYWNrIHsoZnVuY3Rpb24oZXY6IE1lc3NhZ2VFdmVudCkgOiBhbnkpfVxuICogQHJldHVybnMge1Byb21pc2U8V2ViU29ja2V0Pn1cbiAqL1xuZXhwb3J0IGFzeW5jIGZ1bmN0aW9uIGdldFdzKG9uTWVzc2FnZUNhbGxiYWNrKSB7XG5cbiAgICBpZiAod3NDb25uZWN0aW9uUHJvbWlzZSkge1xuICAgICAgICByZXR1cm4gd3NDb25uZWN0aW9uUHJvbWlzZTtcbiAgICB9XG5cblxuICAgIHdzQ29ubmVjdGlvblByb21pc2UgPSBuZXcgUHJvbWlzZSgocmVzb2x2ZSwgcmVqZWN0KSA9PiB7XG4gICAgICAgIGlmICh3c0Nvbm5lY3Rpb24pIHtcbiAgICAgICAgICAgIHJlc29sdmUod3NDb25uZWN0aW9uKVxuICAgICAgICAgICAgcmV0dXJuXG4gICAgICAgIH1cblxuXG4gICAgICAgIGNvbnN0IGNvbm5lY3QgPSBmdW5jdGlvbiAocmVzb2x2ZSkge1xuICAgICAgICAgICAgY29uc3Qgd3NIb3N0ID0gd2luZG93LmxvY2F0aW9uLmhvc3RuYW1lICsgKHdpbmRvdy5sb2NhdGlvbi5wb3J0ID8gXCI6XCIgKyB3aW5kb3cubG9jYXRpb24ucG9ydCA6IFwiXCIpXG4gICAgICAgICAgICBjb25zdCBwcm90b2NvbCA9IHdpbmRvdy5sb2NhdGlvbi5wcm90b2NvbCA9PSBcImh0dHA6XCIgPyBcIndzXCIgOiBcIndzc1wiXG5cbiAgICAgICAgICAgIGNvbnN0IHdzVXJsID0gcHJvdG9jb2wgKyBcIjovL1wiICsgd3NIb3N0ICsgXCIvdmlkZW8vcm9vbS93c1wiXG5cbiAgICAgICAgICAgIGNvbnN0IGxvY2FsV3NDb25uZWN0aW9uID0gbmV3IFdlYlNvY2tldCh3c1VybCk7XG4gICAgICAgICAgICBsb2NhbFdzQ29ubmVjdGlvbi5vbm1lc3NhZ2UgPSBvbk1lc3NhZ2VDYWxsYmFja1xuXG5cbiAgICAgICAgICAgIGxvY2FsV3NDb25uZWN0aW9uLm9uY2xvc2UgPSBhc3luYyBmdW5jdGlvbiAoZSkge1xuICAgICAgICAgICAgICAgIHdzQ29ubmVjdGlvblByb21pc2UgPSB1bmRlZmluZWRcbiAgICAgICAgICAgICAgICB3c0Nvbm5lY3Rpb24gPSB1bmRlZmluZWRcbiAgICAgICAgICAgICAgICBzZXRUaW1lb3V0KGNvbm5lY3QsIDEwMDApXG4gICAgICAgICAgICB9XG5cbiAgICAgICAgICAgIGxvY2FsV3NDb25uZWN0aW9uLm9ub3BlbiA9IGFzeW5jIGZ1bmN0aW9uIChlKSB7XG4gICAgICAgICAgICAgICAgd3NDb25uZWN0aW9uID0gbG9jYWxXc0Nvbm5lY3Rpb25cbiAgICAgICAgICAgICAgICBpZiAocmVzb2x2ZSkge1xuICAgICAgICAgICAgICAgICAgICByZXNvbHZlKGxvY2FsV3NDb25uZWN0aW9uKVxuICAgICAgICAgICAgICAgIH1cbiAgICAgICAgICAgIH1cbiAgICAgICAgfVxuXG4gICAgICAgIGNvbm5lY3QocmVzb2x2ZSlcbiAgICB9KVxuXG4gICAgcmV0dXJuIHdzQ29ubmVjdGlvblByb21pc2Vcbn0iXSwibmFtZXMiOltdLCJzb3VyY2VSb290IjoiIn0=