/* global ipfs */

const run = async (files) => {
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
	await ipfs.files.mkdir('/some/stuff', { parents: true })
	const rootDirectoryContents = await ipfs.files.ls('/', { long: true })
	const filepathsToMove = rootDirectoryContents.filter(file => file.type === 0).map(file => '/' + file.name)
	await ipfs.files.mv(filepathsToMove, '/some/stuff')
  
	await ipfs.files.cp('/ipfs/QmWCscor6qWPdx53zEQmZvQvuWQYxx1ARRCXwYVE4s9wzJ', '/some/stuff/success.txt')
  
	let someStuffDirectoryContents = await ipfs.files.ls('/some/stuff', { long: true })
	return someStuffDirectoryContents
}
  
return run
  