/* globals ipfs */

const run = async () => {
	let cid = await ipfs.dag.put({ test: 1 })
	let cid2 = await ipfs.dag.put({ bar: cid })
	let node = await ipfs.dag.get(cid2, '/bar/test')
	return node.value
}
  
return run
  