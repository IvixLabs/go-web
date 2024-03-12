"use strict";
/*
 * ATTENTION: The "eval" devtool has been used (maybe by default in mode: "development").
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
(self["webpackChunk"] = self["webpackChunk"] || []).push([["video"],{

/***/ "./video.js":
/*!******************!*\
  !*** ./video.js ***!
  \******************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var _shared__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./shared */ \"./shared.js\");\n/* harmony import */ var _video_peer__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./video/peer */ \"./video/peer.js\");\n\n\n\n\nfunction main() {\n    const uuid = document.getElementById(\"uuid\").value\n\n    const startButton = document.getElementById('startButton')\n    const recordButton = document.getElementById(\"recordButton\")\n    const stopButton = document.getElementById(\"stopButton\")\n\n    let devices\n\n    const localVideo = document.getElementById('localVideo')\n    const remoteVideo = document.getElementById('remoteVideo')\n    const resolutionSelect = document.getElementById(\"selectResolution\")\n    const cameraSelect = document.getElementById(\"selectCamera\")\n    const msgInput = document.getElementById(\"msgInput\")\n    const msgButton = document.getElementById(\"msgButton\")\n    const msgChat = document.getElementById(\"msgChat\")\n\n    const localPeer = (0,_video_peer__WEBPACK_IMPORTED_MODULE_1__.createPeer)(remoteVideo, uuid)\n\n    const dataChannel = localPeer.createDataChannel(\"messages\",\n        {negotiated: false}\n    )\n\n    dataChannel.onmessage = function (event) {\n        console.log(event)\n\n        let newParagraph = document.createElement(\"p\");\n        let textNode = document.createTextNode(event.data);\n        newParagraph.appendChild(textNode);\n        msgChat.appendChild(newParagraph)\n    }\n\n    msgButton.onclick = function () {\n        const msg = msgInput.value\n        console.log('send msg')\n        console.log(msg)\n        dataChannel.send(msg)\n    }\n\n    recordButton.onclick = async function () {\n        return captureCamera(localPeer)\n    }\n\n    stopButton.onclick = async function () {\n        return stopCamera(localPeer)\n    }\n\n    let hasMicrophone = true;\n\n    startButton.onclick = async function () {\n        startButton.disabled = true\n        await listDevices()\n        recordButton.classList.remove(\"d-none\")\n        stopButton.classList.remove(\"d-none\")\n        startButton.classList.add(\"d-none\")\n\n    }\n\n\n    async function askForPermissions() {\n        let stream;\n        try {\n            const constraints = {video: true, audio: true}\n            stream = await navigator.mediaDevices.getUserMedia(constraints)\n        } catch (error) {\n            console.log(error)\n            const constraints = {video: true, audio: false}\n            stream = await navigator.mediaDevices.getUserMedia(constraints)\n            hasMicrophone = false\n        }\n        closeStream(stream)\n    }\n\n    function closeStream(stream) {\n        try {\n            if (stream) {\n                stream.getTracks().forEach(track => track.stop());\n            }\n        } catch (e) {\n            alert(e.message);\n        }\n    }\n\n    async function listDevices() {\n        const devices = await getCameraDevices()\n        for (let index = 0; index < devices.length; index++) {\n            const device = devices[index];\n\n            const option = document.createElement(\"option\");\n            option.value = device.deviceId\n            option.text = device.label !== \"\" ? device.label : \"Camera \" + index\n            cameraSelect.add(option);\n        }\n    }\n\n    async function getCameraDevices() {\n        await askForPermissions()\n        devices = await navigator.mediaDevices.enumerateDevices()\n        const cameraDevices = []\n        for (let i = 0; i < devices.length; i++) {\n            const device = devices[i]\n            if (device.kind === 'videoinput') {\n                cameraDevices.push(device)\n            }\n        }\n        return cameraDevices;\n    }\n\n\n    /**\n     * @param peer {RTCPeerConnection}\n     * @returns {Promise<void>}\n     */\n    async function stopCamera(peer) {\n        peer.getSenders().forEach(function (sender) {\n            sender.track.stop()\n        })\n    }\n\n    /**\n     * @param peer {RTCPeerConnection}\n     * @returns {Promise<void>}\n     */\n    async function captureCamera(peer) {\n        let constraints = {\n            audio: false,\n            video: true\n        }\n\n        if (cameraSelect.selectedOptions[0]) {\n            let device = devices[cameraSelect.selectedIndex];\n            let deviceID = device.deviceId;\n\n            constraints = {\n                video: {deviceId: deviceID},\n                audio: hasMicrophone\n            }\n        }\n        if (resolutionSelect.selectedOptions[0]) {\n            let desiredWidth = resolutionSelect.selectedOptions[0].label.split(\"x\")[0];\n            let desiredHeight = resolutionSelect.selectedOptions[0].label.split(\"x\")[1];\n            if (constraints[\"video\"][\"deviceId\"]) {\n                constraints[\"video\"][\"width\"] = desiredWidth;\n                constraints[\"video\"][\"height\"] = desiredHeight;\n            } else {\n                constraints = {width: desiredWidth, height: desiredHeight};\n            }\n        }\n\n        navigator.mediaDevices.getUserMedia(constraints).then(function (stream) {\n\n            stopCamera(peer)\n\n            console.log(peer.getSenders())\n            peer.getSenders()\n\n            let foundSender = false\n            peer.getSenders().forEach(function (sender) {\n                sender.setStreams(stream)\n                //     if (sender.track.kind === track.kind) {\n                //         sender.replaceTrack(track)\n                foundSender = true\n                //     }\n            })\n\n            if (!foundSender) {\n                for (const track of stream.getTracks()) {\n                    peer.addTrack(track, stream)\n                }\n            } else {\n                for (const track of stream.getTracks()) {\n                    peer.getSenders().forEach(function (sender) {\n                        if (sender.track.kind === track.kind) {\n                            sender.replaceTrack(track)\n                        }\n                    })\n                }\n            }\n\n            // for (const track of stream.getTracks()) {\n\n            // let foundSender = false\n            // peer.getSenders().forEach(function (sender) {\n            //     if (sender.track.kind === track.kind) {\n            //         sender.replaceTrack(track)\n            //         foundSender = true\n            //     }\n            // })\n\n            // if (!foundSender) {\n            // peer.addTrack(track, stream)\n            // }\n            // }\n\n            localVideo.srcObject = stream;\n\n\n        }).catch(function (error) {\n            alert('Unable to capture your camera. Please check console logs.');\n            console.error(error);\n        });\n    }\n\n}\n\n(function () {\n    main()\n})();\n\n//# sourceURL=webpack:///./video.js?");

/***/ }),

/***/ "./video/peer.js":
/*!***********************!*\
  !*** ./video/peer.js ***!
  \***********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   createPeer: () => (/* binding */ createPeer)\n/* harmony export */ });\n/* harmony import */ var _ws__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./ws */ \"./video/ws.js\");\n\n\n/**\n * @param remoteVideo {HTMLMediaElement}\n * @param uuid {string}\n * @returns {RTCPeerConnection}\n */\nfunction createPeer(remoteVideo, uuid) {\n    const peerConnectionConfig = {\n        iceServers: [\n            {'urls': 'stun:stun.stunprotocol.org:3478'},\n            {'urls': 'stun:stun.l.google.com:19302'},\n        ]\n    }\n\n    const peer = new RTCPeerConnection(peerConnectionConfig)\n\n    peer.onicecandidate = async function (e) {\n        console.log(\"localPeer onicecandidate\", e)\n\n        if (e.candidate != null) {\n            const ws = await (0,_ws__WEBPACK_IMPORTED_MODULE_0__.getWs)(getOnMessageWsCallback(peer, uuid))\n            ws.send(JSON.stringify({ice: e.candidate, uuid: uuid}));\n        }\n    }\n\n    peer.onnegotiationneeded = async function (e) {\n        console.log(\"localPeer onnegotiationneeded\", e)\n        return enterInRoom(this, uuid)\n    }\n\n    /**\n     * @param {RTCTrackEvent} e\n     */\n    peer.ontrack = async function (e) {\n        console.log(\"localPeer ontrack\", e)\n        remoteVideo.srcObject = e.streams[0];\n    }\n\n    return peer\n}\n\n/**\n * @param peer {RTCPeerConnection}\n * @param uuid {string}\n * @returns {Promise<void>}\n */\nasync function enterInRoom(peer, uuid) {\n    if (!peer.localDescription) {\n        const offerDescription = await peer.createOffer()\n        await peer.setLocalDescription(offerDescription)\n    }\n\n    const ws = await (0,_ws__WEBPACK_IMPORTED_MODULE_0__.getWs)(getOnMessageWsCallback(peer, uuid))\n    ws.send(JSON.stringify({sdp: peer.localDescription, uuid: uuid}))\n}\n\n\n/**\n * @param peer {RTCPeerConnection}\n * @param uuid {string}\n * @returns {(function(message : MessageEvent): Promise<void>)}\n */\nfunction getOnMessageWsCallback(peer, uuid) {\n    return async function (message) {\n\n        const signal = JSON.parse(message.data)\n        console.log(\"signal\", signal)\n\n        if (signal.uuid === uuid) {\n            return\n        }\n\n\n        if (signal.sdp) {\n            if (peer.signalingState !== \"stable\") {\n                await peer.setRemoteDescription(signal.sdp)\n\n                if (signal.sdp.type === 'offer') {\n                    const localDescription = await peer.createAnswer()\n                    await peer.setLocalDescription(localDescription)\n                    await enterInRoom(peer, uuid)\n                }\n            }\n        }\n\n        if (signal.ice) {\n            if (peer.remoteDescription) {\n                await peer.addIceCandidate(signal.ice)\n            }\n        }\n    }\n}\n\n//# sourceURL=webpack:///./video/peer.js?");

/***/ }),

/***/ "./video/ws.js":
/*!*********************!*\
  !*** ./video/ws.js ***!
  \*********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony export */ __webpack_require__.d(__webpack_exports__, {\n/* harmony export */   getWs: () => (/* binding */ getWs)\n/* harmony export */ });\nlet wsConnection;\nlet wsConnectionPromise\n\n/**\n *\n * @param onMessageCallback {(function(ev: MessageEvent) : any)}\n * @returns {Promise<WebSocket>}\n */\nasync function getWs(onMessageCallback) {\n\n    if (wsConnectionPromise) {\n        return wsConnectionPromise;\n    }\n\n\n    wsConnectionPromise = new Promise((resolve, reject) => {\n        if (wsConnection) {\n            resolve(wsConnection)\n            return\n        }\n\n\n        const connect = function (resolve) {\n            const wsHost = window.location.hostname + (window.location.port ? \":\" + window.location.port : \"\")\n            const protocol = window.location.protocol == \"http:\" ? \"ws\" : \"wss\"\n\n            const wsUrl = protocol + \"://\" + wsHost + \"/video/room/ws\"\n\n            const localWsConnection = new WebSocket(wsUrl);\n            localWsConnection.onmessage = onMessageCallback\n\n\n            localWsConnection.onclose = async function (e) {\n                wsConnectionPromise = undefined\n                wsConnection = undefined\n                setTimeout(connect, 1000)\n            }\n\n            localWsConnection.onopen = async function (e) {\n                wsConnection = localWsConnection\n                if (resolve) {\n                    resolve(localWsConnection)\n                }\n            }\n        }\n\n        connect(resolve)\n    })\n\n    return wsConnectionPromise\n}\n\n//# sourceURL=webpack:///./video/ws.js?");

/***/ })

},
/******/ __webpack_require__ => { // webpackRuntimeModules
/******/ var __webpack_exec__ = (moduleId) => (__webpack_require__(__webpack_require__.s = moduleId))
/******/ var __webpack_exports__ = (__webpack_exec__("./video.js"));
/******/ }
]);