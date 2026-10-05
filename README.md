# Go scripts

Small Go scripts for automating everyday tasks

## Scripts

### organize-images

Moves images from `~/Downloads` into `~/Downloads/images`

Supported formats:

- `.png`
- `.jpg`
- `.webp`
- `.jpeg`
- `.gif`
- `.heif`

Run:

	go run ./cmd/organize-images

### organize-pdfs

Moves pdfs from `~/Downloads` into `~/Downloads/pdfs`

Run:

	go run ./cmd/organize-pdfs
	
### clean-dmgs

Removes `.dmg` files from `~/Downloads`

Run:

	go run ./cmd/clean-dmgs
	
### clean-nodemodules

Removes `node_modules` and `.next` folders from `~/Work/Projects` when not worked on for more than 7 days

Run:

	go run ./cmd/clean-nodemodules
	
