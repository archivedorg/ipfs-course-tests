/* globals ipfs */

const run = async () => {
	let cid = await ipfs.dag.put({ test: 1 })
	// your code goes here
}
  
return run
  