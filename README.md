<img src="assets/logo.png" width="200">

---

- [API definition](#api-definition)
  - [GET v1/image](#get-v1image)



## API definition


### GET v1/image

Supported params:

- `url`
  -  an HTTP 1.1 or HTTP2 based URL that returns an image when GET HTTP method is used.
  - If an image is not returned from the URL, the transformation will fail with HTTP 400 (BadRequest)
  - Supported formats: `JP[E]G, PNG, WebP, GIF, BMP`
- `format`
  - Converts the provided image to the specified format. Supported `WebP, JPG, JPEG, PNG, GIF`. If the format is not specified, the original format is returned.
- `quality`
  - Image quality when returning the new image.
- `resizeW`
- `resizeH`
- `aspectRatio`
- `filter[<type>]`
  - Multiple filters can be applied at the same time, such as:
    - `filter[grayscale]`:
    - `filter[blur]`:
- `rotateDeg`
  - rotates based on a degree, e.g. `180` rotates 180 degrees.
- `flip=true`
  - Flips the image (inverts it completely)
- `cacheControl=`
  - If specified, will return the value as part of the response headers in order to instruct clients of a particular cache control.

The result will be an image with all the transformations applied. 
A subsequent request will return a cached version of the transformed image (if params are kept the same).