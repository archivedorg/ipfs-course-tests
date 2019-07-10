// Run files.stat to check the status of your root directory ( / ) in IPFS. 
// Be sure to return your result.

/* global ipfs */

const run = async () => {
	return await ipfs.files.stat('/')
}
  
return run
  