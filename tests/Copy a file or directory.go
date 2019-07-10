/* global ipfs */

const run = async (files) => {
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
	await ipfs.files.mkdir('/some/stuff', { parents: true })
	const rootDirectoryContents = await ipfs.files.ls('/', { long: true })
	const filepathsToMove = rootDirectoryContents.filter(file => file.type === 0).map(file => '/' + file.name)
	await ipfs.files.mv(filepathsToMove, '/some/stuff')
  
	// Your code goes here
  
	let someStuffDirectoryContents = await ipfs.files.ls('/some/stuff', { long: true })
	return someStuffDirectoryContents
}
  
return run
  