/* global ipfs */

const run = async (files) => {
	await Promise.all(files.map(f => ipfs.files.write('/' + f.name, f, { create: true })))
	await ipfs.files.mkdir('/some/stuff', { parents: true })
	const rootDirectoryContents = await ipfs.files.ls('/', { long: true })
  
	const filepathsToMove = rootDirectoryContents.filter(file => file.type === 0).map(file => '/' + file.name)
	await ipfs.files.mv(filepathsToMove, '/some/stuff')
  
	//  // alternatively, wrapping multiple mv calls into a single async function with await:
	//  const filesToMove = rootDirectoryContents.filter(item => item.type === 0)
	//  await Promise.all(filesToMove.map(file => {
	//    return ipfs.files.mv('/' + file.name, '/some/stuff')
	// }))
  
	const someStuffDirectoryContents = await ipfs.files.ls('/some/stuff', { long: true })
	return someStuffDirectoryContents
}
  
return run
  