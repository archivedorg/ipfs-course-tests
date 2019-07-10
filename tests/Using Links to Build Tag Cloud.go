/* globals ipfs */

const run = async () => {
	const natCid = await ipfs.dag.put({ author: "Nat" })
	const samCid = await ipfs.dag.put({ author: "Sam" })
	const treePostCid = await ipfs.dag.put({
	  content: "trees",
	  author: samCid,
	  tags: ["outdoor", "hobby"]
	})
	const computerPostCid = await ipfs.dag.put({
	  content: "computers",
	  author: natCid,
	  tags: ["hobby"]
	})
  
	// Add your code here
}
  
return run
  