const WebSocket = require('ws')
const shortid = require('shortid')

const Subscriber = require('./subscriber')

let wsServer

const subscribers = new Subscriber()

exports.start = async function ({port, gatewayClient}) {
  wsServer = new WebSocket.Server({
    port: port
  })

  wsServer.on('connection', (conn) => {
    const connId = shortid.generate()
    conn.connId = connId
    conn.on('message', (raw) => {
      const data = JSON.parse(raw)
      const packet = data.packet
      switch (data.type) {
        case 'publish':
          subscribers.send(packet.topic, packet.payload)
          break
        case 'subscribe':
          subscribers.add(packet.topic, connId, conn)
          break
        case 'unsubscribe':
          subscribers.del(packet.topic, connId)
          break
      }
    })
    conn.on('close', () => {
      subscribers.delAll(connId)
    })
  })
}

exports.stop = async function () {

}
