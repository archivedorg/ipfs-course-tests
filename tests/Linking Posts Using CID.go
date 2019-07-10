/* globals ipfs */

const run = async () => {
	const natCid = await ipfs.dag.put({ author: "Nat" })
	const samCid = await ipfs.dag.put({ author: "Sam" })
  
	// Modify the blog posts below
	const treePostCid = await ipfs.dag.put({ content: "trees" })
	const computerPostCid = await ipfs.dag.put({ content: "computers" })
  
	return [treePostCid, computerPostCid]
}
  
return run
  