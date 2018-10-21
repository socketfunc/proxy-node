const assert = require('power-assert')

const HashRing = require('../../lib/hashring')

describe('HashRing', () => {
  it('#getNode', () => {
    const hosts = ['192.168.10.1', '192.168.10.2', '192.168.10.3', '192.168.10.4']
    const hash = new HashRing(hosts)
    assert('192.168.10.1', hash.getNode('AAA'))
  })
})
