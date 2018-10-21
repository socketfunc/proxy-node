const crypto = require('crypto')
const events = require('events')

class HashRing extends events.EventEmitter {
  constructor (addrs, vnode) {
    super()

    this.vnode = vnode || 100
    this.nodes = this.createNodes(addrs)
    this.ring = this.createRing(addrs)
    this.ringKeys = Object.keys(this.ring).sort()
  }

  update (addrs) {
    const newNodes = this.createNodes(addrs)
    const diff = this.diff(newNodes, this.nodes)
    if (diff.join.length > 0 || diff.leave.length > 0) {
      this.nodes = this.createNodes(diff.keep.concat(diff.join))
      this.ring = this.createRing(Object.keys(this.nodes))
      this.ringKeys = Object.keys(this.ring).sort()
      this.emit('change', diff)
    }
  }

  getNode (key) {
    const hash = this.hash(key)
    const lastIdx = this.ringKeys.length - 1
    let head = 0
    let tail = lastIdx

    while (head <= tail) {
      const pos = head + Math.floor((tail - head) / 2)
      const ph = this.ringKeys[pos]
      if (hash === ph) {
        return this.ring[ph]
      } else if (hash < ph) {
        tail = pos - 1
      } else {
        head = pos + 1
      }
    }

    if (head > lastIdx) {
      return this.ring[this.ringKeys[0]]
    }
    return this.ring[this.ringKeys[head]]
  }

  hash (key) {
    return crypto.createHash('md5').update(key).digest('hex')
  }

  createRing (addrs) {
    const ring = {}
    addrs.forEach((addr) => {
      for (let i = 0; i < this.vnode; i++) {
        const h = this.hash(`${addr}-${i}`)
        ring[h] = addr
      }
    })
    return ring
  }

  createNodes (addrs) {
    return addrs.reduce((obj, val) => {
      obj[val] = true
      return obj
    }, {})
  }

  diff (newNodes, oldNodes) {
    const diff = {
      join: [],
      leave: [],
      keep: []
    }

    Object.keys(newNodes).forEach((node) => {
      if (!oldNodes[node]) {
        diff.join.push(node)
      }
    })

    Object.keys(oldNodes).forEach((node) => {
      if (newNodes[node]) {
        diff.keep.push(node)
      } else {
        diff.leave.push(node)
      }
    })

    return diff
  }
}

module.exports = HashRing
