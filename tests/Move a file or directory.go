/* global ipfs */

const run = async (files) => {
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
	await ipfs.files.mkdir('/some/stuff', { parents: true })
	const rootDirectoryContents = await ipfs.files.ls('/', { long: true })
  
	const filesToMove = // create an array of files to be moved (no directories)
  
	const filepathsToMove = // create an array of the paths of those files
  
	// move all the files in filepathsToMove into /some/stuff
  
	const someStuffDirectoryContents = await ipfs.files.ls('/some/stuff', { long: true })
	return someStuffDirectoryContents
}
  
return run
  