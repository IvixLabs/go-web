import {getWs} from "./ws";

/**
 * @param remoteVideo {HTMLMediaElement}
 * @param uuid {string}
 * @returns {RTCPeerConnection}
 */
export function createPeer(remoteVideo, uuid) {
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
            const ws = await getWs(getOnMessageWsCallback(peer, uuid))
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

    const ws = await getWs(getOnMessageWsCallback(peer, uuid))
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