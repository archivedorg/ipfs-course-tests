const run = async (files) => {
	// This code adds your uploaded files to your root directory in IPFS
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
  
	await ipfs.files.mkdir('/some/stuff', { parents: true })
  
	let rootDirectoryContents = await ipfs.files.ls('/', { long: true })
	return rootDirectoryContents
}
  
return run
  