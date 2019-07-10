/* globals ipfs */

const run = async () => {
	let cid = await ipfs.dag.put({ test: 1 })
	let cid2 = await ipfs.dag.put({ bar: cid })
	return cid2
}
  
return run
  