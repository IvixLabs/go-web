let wsConnection;
let wsConnectionPromise

/**
 *
 * @param onMessageCallback {(function(ev: MessageEvent) : any)}
 * @returns {Promise<WebSocket>}
 */
export async function getWs(onMessageCallback) {

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