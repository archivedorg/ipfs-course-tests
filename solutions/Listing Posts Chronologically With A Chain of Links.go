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
	  tags: ["hobby"],
	  prev: treePostCid
	})
	const dogPostCid = await ipfs.dag.put({
	  content: "dogs",
	  author: samCid,
	  tags: ["funny", "hobby"],
	  prev: computerPostCid
	})
  
	const outdoorTagCid = await ipfs.dag.put({
	  tag: "outdoor",
	  posts: [treePostCid]
	})
	const hobbyTagCid = await ipfs.dag.put({
	  tag: "hobby",
	  posts: [treePostCid, computerPostCid, dogPostCid]
	})
	const funnyTagCid = await ipfs.dag.put({
	  tag: "funny",
	  posts: [dogPostCid]
	})
  
	return dogPostCid
}
  
return run
  