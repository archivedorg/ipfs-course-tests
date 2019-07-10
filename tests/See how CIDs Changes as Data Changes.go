/* global ipfs */

const run = async (files) => {
	// this code adds your uploaded files to IPFS
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
	const rootDirectoryContents = await ipfs.files.ls('/', { long: true })
  
	const directoryStatus = // your code goes here
  
	return directoryStatus
}
  
return run
  