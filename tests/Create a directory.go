const run = async (files) => {
	// This code adds your uploaded files to your root directory in IPFS
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
  
	// Add your code to create a new directory here
  
	let rootDirectoryContents = await ipfs.files.ls('/', { long: true })
	return rootDirectoryContents
}

return run
  