/* globals ipfs */

const run = async () => {
	let cid = await ipfs.dag.put({ test: 1 })
	return cid
}
  
return run
  