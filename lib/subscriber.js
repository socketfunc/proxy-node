const WebSocket = require('ws')

class Subscriber {
  constructor () {
    this.conns = {}
    this.topics = {}
  }

  add (topic, id, conn) {
    if (this.conns[id]) {
      return
    }

    if (!Array.isArray(this.topics[topic])) {
      this.topics[topic] = []
    }
    this.topics[topic].push(id)
    this.conns[id] = conn
  }

  del (topic, id) {
    const ids = this.topics[topic]
    ids.forEach((connId, index) => {
      if (id === connId) {
        ids.splice(index, 1)
      }
    })
    delete this.conns[id]
  }

  delAll (id) {
    Object.keys(this.topics).forEach((topic) => {
      this.del(topic, id)
    })
  }

  send (topic, payload) {
    const ids = this.topics[topic]
    ids.forEach((id) => {
      const conn = this.conns[id]
      if (conn.readyState === WebSocket.OPEN) {
        conn.send(payload)
      }
    })
  }
}

module.exports = Subscriber