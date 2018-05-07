package heraldry

import svg "github.com/ajstarks/svgo"

func renderCrownToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M99.1,124.7c2.4-0.6,4.8-1.1,7-2.1c0.4-0.2,0.8-0.3,1.2-0.4c1.4-0.3,2.3,0.3,2.6,1.7c0.1,0.6,0.1,1.3,0,1.9    c-0.3,2.6,0.5,4,2.9,5c1.6,0.7,3.3,1,4.8,1.8c1.7,0.9,1.9,1.9,0.8,3.4c-0.6,0.8-1.4,1.5-2.2,2.2c-1,0.9-1.9,1.8-2.6,2.9    c-0.8,1.4-0.9,2.9-0.1,4.3c0.2,0.4,0.5,0.8,0.8,1.2c1.1,1.6,0.9,2.4-0.9,3.2c-3.2,1.5-6.5,1.6-9.9,0.4c-0.6-0.2-1.2-0.5-1.8-0.8    c-0.5-0.2-1-0.4-1.5-0.5c-0.8-0.2-1.6,0-2.1,0.7c-0.5,0.7-0.1,1.3,0.3,1.9c2,3.2,3.3,6.6,4.5,10.1c1.2,3.6,2.3,7.2,4.1,10.6    c0.8,1.5,1.8,3,2.5,4.6c0.6,1.4,1.6,2.5,3,3.2c2.1,1.2,4.2,0.9,5.8-1c1.3-1.5,2-3.3,2.6-5.1c1.6-4.7,2.7-9.6,4-14.4    c0.1-0.3,0.1-0.7,0.1-1.1c0-0.5-0.1-1,0.5-1.3c0.6-0.2,1.2-0.1,1.7,0.3c0.5,0.4,0.8,0.9,0.8,1.5c0.1,3.7,1.7,6.9,3.4,10.1    c0.8,1.6,1.5,3.3,2.4,4.9c1.4,2.6,3.3,3.4,5.7,2.6c1.6-0.5,2.5-1.7,3-3.2c1.1-3.4,2.1-6.8,3.1-10.2c1.7-6.1,3.3-12.2,4.2-18.4    c0.1-0.6,0.2-1.1,0.1-1.7c-0.1-0.6-0.5-1.1-1-1.3c-0.4-0.1-0.6,0.4-0.8,0.7c-3.7,4.2-8.3,5.2-13.6,4.2c-1.3-0.3-1.6-0.8-1.3-2.1    c0.1-0.4,0.2-0.8,0.4-1.1c0.7-1.9,0.5-2.9-0.9-4.3c-0.8-0.8-1.9-1.4-2.8-2c-1.4-0.9-2.8-1.8-4-3.1c-1.3-1.4-1.2-2.6,0.2-3.8    c1.3-1.1,2.8-1.9,4.3-2.7c1.2-0.6,2.5-1.2,3.7-2c1.3-1,1.6-2,0.7-3.4c-0.5-0.8-1-1.4-1.7-2c-0.6-0.4-1-0.9-0.8-1.7    c0.2-0.8,0.7-1.4,1.6-1.6c1.7-0.4,3.4-0.5,5.1,0.2c0.7,0.3,1.3,0.4,2.2,0.7c-0.4-2.4-1.1-4.5-2.1-6.4c-0.4-0.8-0.8-1.5-0.8-2.3    c-0.1-1.7,0.9-2.5,2.5-2.2c3.1,0.7,4.2,0,5.4-3c0.9-2.3,1.8-4.7,2.5-7.1c0.4-1.6,1.2-3,3-3.2c1.9-0.3,3.8,0.1,5,1.7    c1.7,2.3,2.7,4.9,2.6,7.8c-0.1,3,1.1,4.1,4.1,3.6c0.7-0.1,1.3-0.3,2-0.3c1.9,0,2.5,0.9,1.7,2.6c-0.8,1.8-1.2,3.8-1.5,5.7    c-0.4,2.6,0.8,3.7,3.4,3.3c0.8-0.1,1.6-0.3,2.5-0.5c2.9-0.7,4.7,1.2,4,4.1c0,0.2-0.1,0.3-0.2,0.5c-1.2,3.5-1,4.3,2.1,6.4    c0.9,0.6,1.8,1.2,2.5,2c1.3,1.4,1.2,2.7-0.1,4.1c-0.8,0.8-1.7,1.3-2.6,1.9c-0.9,0.6-1.8,1.1-2.6,1.7c-2,1.5-2.9,3.5-2.4,6    c0.5,2.7-1.1,4-3.7,2.8c-3-1.4-5.9-2.7-8.7-4.6c-0.8-0.6-1.7-1.1-2.8-1c-0.7,0-1.5,0.3-1.8,0.9c-0.4,0.7,0.2,1.1,0.5,1.6    c0.8,1,0.9,2.2,1.2,3.3c1.8,6.9,3.2,13.9,5.8,20.6c0.8,2.2,1.8,4.4,3,6.5c2.2,4,4.7,4.2,7.5,0.7c2.3-2.8,3.6-6.1,4.9-9.4    c0.7-1.8,1.3-3.5,1.9-5.3c0.3-1,0.8-1.7,2-1.5c1.1,0.2,1.3,1,1.3,2c0,4-0.2,8.1,0.4,12.1c0.2,1.2,0.4,2.4,0.9,3.4    c0.3,0.6,0.7,1.2,1.1,1.7c1.8,1.8,3.2,1.6,4.6-0.5c3.9-5.8,7.3-12,10.5-18.2c1.1-2.1,2.2-4.2,3.3-6.4c0.4-0.8,0.8-1.5,1-2.4    c0.1-0.8-0.2-1.4-0.8-1.9c-0.6-0.4-1.1-0.2-1.5,0.2c-2.5,2-5.3,1.9-8.2,1.3c-2.4-0.5-3-1.7-1.7-3.8c0.2-0.4,0.5-0.8,0.7-1.3    c1-2.1,0.5-3.6-1.6-4.7c-2.2-1.1-2.9-3.4-3.8-5.4c-0.5-1.1,0.5-1.7,1.5-2.1c1.2-0.4,2.4-0.9,3.5-1.5c2.8-1.6,4.5-3.9,4.7-7.2    c0.1-1.9,0.8-2.3,2.6-1.8c1.2,0.3,2.3,0.8,3.2,1.5c0.9,0.7,1.9,1,2.9,1.3c1.4,0.3,2.1-0.2,2.3-1.6c0.2-1.7,0-3.4-0.2-5.1    c-0.3-2.9,0.5-5.5,2.2-7.7c1.2-1.5,2-1.7,3.7-0.9c2,1,3.3,0.7,4.2-1.4c1-2.2,2.5-4.1,4.2-5.9c0.8-0.8,1.7-1.4,3-1.2    c1.7,0.3,2.7,1.2,2.8,3c0.1,1.8,0.1,3.7,0.1,5.5c0,2.2,0.7,3.1,2.7,3.9c2.1,0.8,2.6,1.6,2.1,3.8c-0.9,3.6-2.1,7.1-5.3,9.5    c-1.2,0.9-2.1,2-2.7,3.3c-0.3,0.5-0.2,0.8,0.4,0.9c0.9,0.2,1.8-0.1,2.7-0.3c2.4-0.5,4,0.1,5.4,2c1.1,1.5,0.8,3-0.7,4.2    c-0.5,0.4-1.1,0.7-1.6,1.1c-1,0.8-1,1.5-0.1,2.4c0.8,0.9,1.7,1.7,2.5,2.6c0.4,0.5,0.8,1.1,1.1,1.7c0.6,1.2,0.2,2.1-1,2.8    c-1.3,0.7-2.8,0.9-4.2,1.2c-0.5,0.1-1,0.2-1.5,0.3c-2.2,0.6-3.1,1.9-2.8,4.2c0.1,0.6,0.2,1.3,0.2,1.9c0,1.7-1,2.6-2.5,2    c-3-1.2-5.9-2.8-7.3-5.9c-0.4-0.9-1-1.5-1.9-1.9c-0.6-0.2-1.2-0.1-1.2,0.6c0,2.6-1.3,4.9-1.7,7.4c-1.2,6.3-1.4,12.6-1.2,19    c0,1.2,0,2.5,0.3,3.7c0.6,2.5,2.6,3.6,4.9,2.7c0.2-0.1,0.3-0.2,0.4-0.4c1.2-2.4,3-4.4,4.7-6.5c1.1-1.3,2.2-2.7,3.1-4.1    c0.5-0.9,2.1-1.3,3-0.8c0.8,0.4,1.1,1.5,0.5,2.6c-1.2,2.5-0.8,5.2-0.7,7.8c0,0.8,0,1.6-0.1,2.4c-0.1,0.8-0.3,1.6,0.5,2.1    c0.9,0.6,1.9,0.7,3,0.2c0.2-0.1,0.5-0.3,0.7-0.5c3.5-3.9,7.6-7.2,11.1-11.1c0.5-0.5,1-1.1,1.4-1.7c1.3-1.9,2.8-3.6,4.7-4.9    c0.9-0.7,1.6-1.6,1.6-2.9c-0.1-1.3-0.7-2-2-2c-0.6,0-1.3,0.1-1.9,0.1c-0.9,0-1.8,0.1-2.5-0.5c-1.8-1.2-2.5-2.7-0.5-4.3    c1.6-1.3,1.7-2.8,0.7-4.8c-0.6-1.1-1.2-2.1-1.7-3.3c-0.7-1.5-0.2-2.3,1.4-2.6c1.4-0.3,2.8-0.4,4.2-0.3c1.2,0.1,1.7-0.4,1.7-1.6    c0-1-0.1-2.1,0-3.1c0-1.4,0.8-2.2,2.2-2.3c0.8,0,1.6,0,2.3,0.4c2,1.1,4.2,2,5.1,4.4c0.3,0.7,0.8,1.1,1.6,0.9    c1-0.3,1.6-0.9,1.7-1.9c0.1-1.2,0-2.4,0.2-3.6c0.7-3.8,3.8-7,7.5-7.8c1.5-0.3,2.5,0.2,2.9,1.8c0.2,0.8,0.3,1.7,0.4,2.6    c0.1,0.4,0,0.8,0.1,1.2c0.2,0.8,0.2,1.9,1.2,2.2c0.9,0.3,1.5-0.6,2-1.2c5.1-5.9,11.1-5.6,17.4-2.7c0.5,0.3,0.5,0.7,0.4,1.1    c-0.3,1.9-2,3.2-4.2,3.2c-5-0.2-9.3,1.4-13.3,4.3c-2,1.4-3.9,3.1-5.9,4.5c-1.5,1.1-2.5,2.3-2.5,4.3c0,1.4-0.6,2.8-1.2,4.1    c-1.3,2.9-3.5,4.3-6.6,4.4c-2.4,0-4.3,1-5.4,3c-6.3,11-13.6,21.5-18,33.6c-1,2.6-1.7,5.3-2.4,8c-0.6,2.5,0.1,4.4,2.2,5.8    c1.2,0.8,2.3,1.6,3.2,2.7c3.1,3.4,3,7.3-0.2,10.6c-1.1,1.1-2.4,2-3.7,2.8c-0.3,0.2-0.7,0.4-1,0.6c-3.3,1.8-3.7,3.4-1.6,6.6    c2.3,3.5,2.3,6-0.7,8.6c-4.6,4.1-3.1,9.6,0.6,12.7c1.1,0.9,2.4,1.8,3.3,2.9c2.4,2.9,1.9,5.9-1.3,7.9c-5.3,3.3-11.1,5.2-17.3,5.9    c-14.5,1.6-28.9,2.8-43.5,3.4c-12.5,0.4-25,0.2-37.5-0.4c-7.9-0.4-15.9-0.9-23.8-1c-12.5-0.2-24.6-2.3-36.5-5.9    c-1.5-0.4-2.9-0.7-4.4-1.1c-2.6-0.7-4.3-2.4-4.9-5c-0.7-2.7-0.1-5.1,2-7c1.1-1,2.3-1.7,3.5-2.5c1.1-0.7,2.1-1.6,2.8-2.7    c1.2-2.1,0.8-4.3-1.2-5.8c-0.8-0.6-1.6-1-2.4-1.6c-1.6-1.2-1.9-2.6-0.9-4.4c0.2-0.4,0.5-0.7,0.7-1.1c0.2-0.3,0.4-0.5,0.6-0.8    c3.2-4.3,2.6-7.3-2.2-9.7c-2.5-1.2-4.2-3-5.2-5.5c-0.5-1.1-0.5-2.2,0-3.3c0.4-0.9,0.8-1.8,1.2-2.6c0.1-0.2,0.2-0.4,0.3-0.5    c2.1-0.9,2.7-3.6,5.1-4.1c0.5-0.1,0.5-0.7,0.5-1.1c0.2-1.5-0.1-3-0.6-4.5c-2-5.4-4.5-10.6-7.3-15.7c-3.7-6.7-7.3-13.4-11-20.1    c-1.9-3.4-4.5-6.2-8.5-7.3c-0.9-0.2-1.9-0.3-2.8-0.5c-2.7-0.6-4.2-2.7-4-5.4c0.4-5.6-1.5-9.1-6.5-11.3c-4.3-1.9-8.8-3.4-13.5-4    c-1.1-0.1-2.2,0-3.2,0.1c-0.8,0-1.6-0.1-2.2-0.7c-0.8-0.9-0.8-2.1,0.2-2.6c3.5-1.9,7.1-3.9,11.4-3.4c1.9,0.2,3.7,0.9,5.1,2.2    c1.4,1.2,3.1,1.7,4.8,2.1c1.1,0.3,1.1-0.7,1-1.3c-0.2-1.4,0.4-2.4,1.1-3.4c0.6-0.8,1.4-1,2-0.2c1.6,1.9,3.6,3.5,5.2,5.5    c1,1.3,1.6,2.7,1.6,4.4c0,1.6,0.7,1.9,2.2,1.4c1-0.4,1.6-1.1,2.3-1.8c0.9-0.9,1.8-1.8,2.9-2.4c1.8-1.1,3.5-0.1,3.6,2    c0,0.8-0.1,1.7-0.1,2.5c0,1.3,0.5,1.8,1.8,1.9c1.4,0.1,2.9,0,4.3,0.2c1.7,0.3,2.2,1.1,1.6,2.7c-0.5,1.4-1.3,2.5-2.1,3.8    c-1,1.5-0.6,2.3,1.2,2.7c0.7,0.1,1.3,0.1,1.9,0.6c1.8,1.6,1.6,3.4-0.4,4.5c-1.6,0.9-3.3,1.4-5.1,1.5c-0.8,0.1-1.7,0.1-2.1,0.9    c-0.4,0.9,0.2,1.5,0.8,2.1c4.1,4.6,7.6,9.8,11.7,14.5c1.3,1.5,2.7,3,4.1,4.4c1.7,1.6,2.8,1.4,4-0.5c1-1.6,1.5-3.4,1.5-5.2    c-0.1-2.4,0.4-4.9,0-7.3c-0.1-0.7,0.2-1.3,0.9-1.7c0.9-0.4,1.8-0.3,2.6,0.4c1.2,1,1.9,2.4,2.5,3.9c0.8,2.1,2.1,4,3.7,5.6    c1.7,1.7,4.3,1,5.1-1.3c0.3-0.8,0.4-1.5,0.4-2.4c0.1-3.8,0.2-7.6,0.2-11.4c0-4-0.2-7.9-1.2-11.8c-0.4-1.5-0.4-1.5-1.7-0.7    c-1,0.6-1.9,1.4-2.9,2.2c-1.5,1.2-3.3,1.8-5.2,1.7c-1.9-0.1-2.8-1.1-2.9-3c-0.1-3.5-0.9-4.4-4.3-5.2c-1.2-0.3-2.5-0.3-3.7-0.9    c-1.3-0.6-1.6-1.4-1-2.7c0.6-1.3,1.6-2.3,2.5-3.4c2.1-2.6,2.1-3.7-0.3-6c-0.8-0.8-0.8-1.8,0.1-2.9c1-1.2,2.5-1.7,4-2.1    c1.5-0.4,3.1-0.6,4.6-0.6c0.9,0,0.9-0.4,0.7-1c-0.3-1.2-1.2-2-2.1-2.7c-1.2-1-2.2-2.1-2.9-3.4c-1.4-2.5-1-4.8,1.1-6.7    c0.2-0.2,0.4-0.4,0.6-0.5c2.3-1.3,2.6-3.1,1.8-5.5c-0.6-1.8-0.8-3.7-0.5-5.6c0.3-1.7,1.2-2.9,2.9-3.4c1.6-0.5,2.9,0,4,1.2    c1.3,1.3,2.1,2.9,2.7,4.6c1.1,3.5,1.8,4.8,5.8,2.7c1.6-0.8,2.9-0.2,3.4,1.6c0.7,2.4,0.4,4.9,0.5,7.4    C99.2,122,99.1,123.3,99.1,124.7z"
	pathData1 := "M122.8,230.8c-3.5-0.2-6-2-8.5-3.7c-2.1-1.4-4.2-2.6-6.8-2.5c-2.7,0-2.7,0-3.7-2.7c-0.3-0.9-0.6-1.9-1.5-2.8    c0.3,2.3,0.2,4.4,1.3,6.2c0.6,1,0.4,2.1-0.3,3.1c-0.7,1-1.4,1.9-2.2,2.8c-0.2,0.2-0.5,0.5-0.8,0.4c-0.4-0.1-0.3-0.5-0.3-0.8    c-0.1-0.4-0.1-0.8-0.2-1.2c-1.2,2-0.3,5.7,1.7,7c0.5,0.3,1.1,0.6,1.7,0.8c2.4,0.8,3.6,2.5,3.9,5c0.1,1.1,0.3,2.3,0.7,3.4    c0.2,0.4,0.1,0.8-0.3,1c-1,0.6-2.1,1.2-3.4,1.1c-1.2,0-2-0.8-2.1-2.3c-0.1-1.6-0.1-3.1-0.1-4.7c0-0.3,0-0.6,0-1    c-0.1-0.5-0.4-0.9-1-0.9c-0.5,0-0.8,0.4-0.9,1c0,0.7,0.1,1.3,0.5,1.8c0.7,1.2,0.8,2.6,0.5,3.9c-0.2,1.3-0.9,2-1.9,2    c-1.2,0-1.5-0.9-1.7-1.8c-0.6-2.7-1.1-5.3-1.6-8c-0.9-4.7-1.7-9.4-2.1-14.2c-0.2-3,0.4-4.1,3.3-4.8c6.8-1.8,13.7-3.2,20.6-4.2    c11.9-1.7,23.8-2.4,35.8-2.5c6.2,0,12.4-0.2,18.6,0.5c2.4,0.3,4.8,0.6,7.1,1.3c2.3,0.7,4.7,0.5,7,0.8c0.6,0.1,1.1,0.1,1.6,0.3    c1.1,0.3,1.7,1,1.7,2.3c0,1.4-0.4,2.7-0.8,4c-0.3,1-0.6,2-0.3,3.1c0.1,0.3,0,0.6,0.3,0.7c0.4,0.1,0.6-0.2,0.8-0.4    c0.7-0.7,0.8-1.6,0.7-2.5c-0.1-1.5,0-3,0.1-4.5c0.1-0.8,0.4-1.5,1.2-1.8c1-0.3,2-0.4,2.9,0.2c0.8,0.5,0.4,1.3,0.2,2    c-0.3,0.9-0.5,1.7-0.2,2.6c0.1,0.4,0.4,0.8,0.8,0.8c0.4,0.1,0.6-0.4,0.7-0.7c0.3-0.5,0.5-1.1,0.7-1.7c0.2-0.5,0.4-1,0.8-1.4    c0.4-0.5,0.9-1,1.5-0.8c0.4,0.1,0.3,0.9,0.4,1.4c0.2,2.6-0.6,4.2-3.1,5c-3.6,1.2-6.5,3.5-9.7,5.4c-1.1,0.7-1.6,1.9-1.1,2.9    c1.3,2.6,1,5.4,0.7,8.1c-0.2,1.1-0.2,2.3-0.3,3.5c-0.1,1.3-0.3,1.4-1.7,1.2c-3.3-0.5-6.6-1-10-0.2c-0.2,0-0.3,0-0.5,0.1    c-1.5,0.9-2.9,0.6-4.2-0.4c-0.4-0.3-0.9-0.4-1.5-0.4c-11.9-0.2-23.9-0.5-35.8,0.5c-6.4,0.6-12.7,1.5-19,2.7    c-2.3,0.4-3.3-0.2-4.1-2.4c-0.3-0.9-0.5-1.9-0.8-2.9c-0.2-0.6,0-0.7,0.5-0.7c0.8,0,1.5,0,2.3-0.1c2.6-0.2,4.6-1.4,5.8-3.7    c1.3-2.5,3.2-4.4,5.5-5.9C122.3,231.2,122.5,231.1,122.8,230.8z"
	pathData2 := "M98,171.9c-0.2,4.3-0.9,8.3-1.1,12.4c-0.1,3.2-4.2,7.3-7.4,7.8c-1.7,0.3-3-0.4-4.1-1.7    c-0.7-0.8-1.3-1.7-2.1-2.4c-1.1-0.8-1.8-0.7-2.6,0.4c-0.5,0.7-0.7,1.6-1,2.4c-0.6,2.1-1.8,3.8-4.1,4.3c-2.2,0.5-4-0.5-5.4-2.1    c-3.7-4.2-7.3-8.5-10.9-12.8c-0.5-0.6-0.7-1.2-0.6-2c0.1-0.3-0.1-0.8,0.3-1c0.3-0.1,0.6,0.3,0.8,0.4c1.7,1.3,3,2.9,4.2,4.5    c1.9,2.4,3.9,4.8,6.7,6.3c0.8,0.4,1.5,1.2,2.2,1.8c1.2,0.9,2.1,0.9,3.4,0.1c2.3-1.6,3.3-4.4,2.6-7.1c-0.5-1.7-0.4-2,1.1-5.5    c1.3,1.8,2.6,3.6,3.9,5.4c0.7,1,1.4,2,2.3,2.8c1.9,1.7,3.8,1.5,5.3-0.5c1.6-2.1,2.3-4.5,2.7-7c0.7-3.7,0.3-7.3,0.2-11    c-0.1-5,0.3-10-0.1-15.1c0.4,0.2,0.5,0.5,0.6,0.8c2.5,6.4,5,12.8,7.5,19.1c1.4,3.7,3.6,6.9,6.6,9.6c2,1.7,4.1,2,6.4,1    c2.3-1,3.9-2.8,4.9-5.2c1-2.5,2.1-5.1,3.1-7.6c0.1-0.3,0.2-0.6,0.4-0.9c0.7-1.2,1.2-1.2,1.9-0.1c1.3,1.8,2.3,3.8,3.4,5.8    c1.1,1.9,2.6,3.5,4.4,4.7c2.5,1.8,4.7,1.5,6.7-0.9c3-3.5,4.7-7.7,6-12c1.4-4.9,2.2-10,3.4-15.1c0.3-1.2,0.6-2.5,1.3-3.7    c1.3,1.9,2.1,4,2.5,6.2c1.2,6,2.8,11.8,5.2,17.4c1.2,2.7,2.5,5.3,4.5,7.5c2.9,3.1,5.9,3.2,8.9,0.2c1.6-1.6,2.7-3.7,3.9-5.6    c0.7-1.1,1.4-2.2,2.2-3.3c0,1.7,0.1,3.4,0.4,5c0.3,1.6,0.8,3.2,1.6,4.7c2.2,4.2,6.6,4.8,9.6,1.1c3-3.6,5.5-7.6,7.9-11.7    c1.9-3.2,3.6-6.5,5.1-9.9c0.8-1.7,2.2-2.9,3.1-4.6c0.2,1.6,0.4,3.3,0.3,4.9c-0.1,4.6-0.9,9.2-1,13.9c-0.1,2.5,0.1,5,0.5,7.5    c0.3,2.1,1.5,3.9,2.9,5.4c0.9,0.9,1.8,1,2.9,0.3c0.6-0.4,1.2-0.9,1.7-1.5c0.7-0.8,1.4-1.7,2.7-2c1-0.2,1.5-1.3,2.1-2    c0.7-0.9,1.4-1.7,2.6-1.9c0.6-0.1,0.8,0,0.9,0.7c0.2,1.4,0,2.9,0,4.3c0,1.9,0.2,3.7,2.1,4.7c2,1.1,4,0.8,5.9-0.4    c1.5-1,2.7-2.3,4-3.4c0.6-0.6,1.3-1.1,2.1-1.9c-1,3.5-2.9,6-5.8,7.6c-1.5,0.9-3.1,1.4-4.8,1.7c-2,0.4-3.2-0.3-3.9-2.2    c-0.5-1.3-0.9-2.6-0.7-4c0.1-0.6,0.1-1.2-0.6-1.5c-0.7-0.3-1.4-0.3-2,0.2c-0.7,0.6-1.3,1.2-1.7,2c-2.2,4.2-9.4,4-12.1-0.4    c-2-3.4-3-7.1-2.4-11.1c0.1-0.7,0.2-1.4,0.1-2.1c-0.2-0.8-0.6-1.3-1.4-1.4c-0.7-0.1-1.2,0.3-1.5,0.9c-0.2,0.3-0.3,0.6-0.4,0.9    c-1.1,4-4,6.9-6.6,10c-1.8,2.1-4.3,3-6.8,3.7c-1.5,0.4-2.8-0.1-3.9-1.1c-1.1-1.1-2.1-2.3-3-3.6c-1.6-2.4-2.5-2.5-4.7-0.6    c-1.1,1-2.2,1.9-3.5,2.5c-3,1.5-5.7,1.1-8.1-1.2c-2.8-2.6-4.6-5.9-5.9-9.5c-1.3-3.5-2.4-7.1-3.6-10.6c-0.2-0.5-0.4-0.9-0.8-1.4    c-0.6,2.2-1.2,4.4-1.8,6.6c-0.4,1.4-0.7,2.9-1.1,4.3c-1.2,4.6-4,8.2-7.5,11.3c-2.1,1.8-4.5,1.9-6.9,0.7c-2.1-1-3.7-2.7-5.2-4.4    c-0.7-0.8-1.3-1.6-2.1-2.3c-0.7-0.6-1.4-0.5-2,0.2c-0.6,0.6-1,1.3-1.4,1.9c-0.7,1.2-1.4,2.3-2.2,3.4c-3.3,4.1-7.3,5.1-12,2.8    c-1.8-0.9-3.4-2.1-4.8-3.6c-2.2-2.6-4.2-5.2-5-8.6C99.6,173.6,98.9,172.7,98,171.9z"
	pathData3 := "M183.6,192.2c-1.3,1.4-1.8,2.9-1.6,4.8c0.4,3-0.9,4.3-3.9,4.1c-0.2,0-0.3,0-0.5,0c-1.8-0.2-1.8-0.2-1.2-1.9    c0.6-1.7,0.8-3.5,0.2-5.6c-0.4,1.3-0.6,2.4-1,3.4c-1,3.2-2,4-5.4,4.1c-0.9,0-1.8,0-2.8,0c-10.3,0-20.6,0.1-30.9-0.1    c-4.4-0.1-8.7,0.6-13,1.2c-0.6,0.1-1.2,0.1-1.8,0.1c-2.8,0-4.2-1.2-4.6-4c-0.1-0.7-0.2-1.4-0.3-2.1c-0.2-1-0.6-1.4-1.7-1.8    c0.7,2,1.2,4,0.9,6.1c-0.2,1.2-0.6,2.1-1.9,2.5c-1.2,0.3-2.1-0.1-2.9-1.3c-1.1-1.6-1.4-3.4-1.4-5.2c-0.1-1.8,0-3.6-0.7-5.3    c-0.2-0.6,0.1-1,0.7-1.1c0.7-0.2,1.5-0.1,2.1,0.2c4.2,1.7,8,0.8,11.5-1.7c1.4-1,2.7-2.1,3.9-3.3c0.8-0.8,1.5-0.8,2.3-0.2    c0.6,0.5,1.3,1,1.9,1.4c5.1,3.7,9.4,3.7,15-0.6c2.8-2.2,4.4-5.3,5.7-8.5c0.4-1,0.7-1.9,1.1-3c1.5,0.8,2.2,2.1,2.8,3.6    c1.8,4.9,5,8.6,9.6,11c1.1,0.6,2.3,1,3.6,1.3c1.3,0.3,2.4,0,3.5-0.8c2.9-2,5.3-1.9,8.2,0.2c1,0.7,2,1.5,2.7,2.5    C183.7,192,183.7,192,183.6,192.2C183.7,192.1,183.6,192.2,183.6,192.2z"
	pathData4 := "M222,212.1c-3.5-0.3-7-1.1-10.5-1.8c-6.2-1.2-12.4-1.8-18.6-2.6c-4.4-0.5-8.7-1-13.1-1.1    c-1.3-0.1-2.5,0.1-3.8,0.5c1.8,0.3,3.6,0.6,5.4,0.9c7.1,1.1,14.4,1.7,21.5,2.9c8.3,1.4,16.3,3.5,23.4,8.2c1.3,0.8,2.5,1.7,3.7,2.7    c0,0.2,0,0.4-0.3,0.2c-4.5-3.8-10.2-4.3-15.5-5.7c-11.9-3.3-24.1-5-36.3-6.1c-12.3-1.1-24.6-1.4-36.9-1.1    c-7.4,0.2-14.6,1-21.9,2.2c-5.7,0.9-11.5,1.6-17.2,2.3c-7.1,0.9-14,2.3-20.8,4.5c-1.8,0.6-3.6,1-5.5,1.3c3.3-3.4,8.1-3.8,12.3-5.9    c-1.8-0.5-3.4-0.5-5-0.2c-2.1,0.4-4.1,1.3-6,2.4c-2.4,1.4-5,2.4-7.5,3.8c-0.4,0.2-0.8,0.3-1.3,0.1c1.2-1.6,2.8-2.8,3.9-4.4    c0.2-0.3,0.7-0.4,1-0.5c9.8-3.5,19.7-6.1,30.1-7.2c2.3-0.3,2.8-1.1,2.3-3.4c-0.9-3.6-1.1-7.3-1.7-11c-0.3-1.8-0.8-3.6-1.6-5.5    c-0.7,2.7-0.3,5.1,0.3,7.5c0.4,1.5,0.8,3,0.9,4.6c0.1,1.3-0.1,2.5-0.6,3.7c-0.4,0.9-1.1,1.1-1.9,0.4c-0.7-0.5-1-1.2-1-2.1    c0-1.5,0-2.9,0-4.4c0-0.9-0.2-1.7-1-2.3c-0.6-0.4-0.9-0.5-1,0.4c-0.2,1.7,0.1,3.3,0.7,4.8c0.7,1.9,0.3,3.7-0.9,5.3    c-0.1,0.2-0.3,0.3-0.5,0.5c-1.2-2.7-2.2-5.3-2.3-8.3c0-1.7,0.5-3.1,1.6-4.4c2-2.3,3.7-4.9,4.9-7.8c0.3-0.8,0.6-1.9,1.6-1.9    c1-0.1,1.5,1,1.9,1.8c2.4,3.9,3.5,8.3,3.7,12.9c0.1,2.4,0.1,4.9-0.2,7.3c-0.1,0.8,0.2,0.8,0.8,0.8c5.4-0.7,10.8-1.3,16.2-1.7    c6.3-0.5,12.6-0.9,18.9-1.2c5.8-0.2,11.6-0.3,17.4-0.1c10.5,0.2,21,1,31.4,2.5c9.7,1.4,19.3,3.4,28.8,6.1c0.3,0.1,0.6,0.2,0.9,0.3    C222.6,212.4,222.3,212.2,222,212.1z"
	pathData5 := "M96.9,255.1c-5.5-0.7-10.6,1-15.6,2.7c-2.1,0.7-4.2,1.4-6.4,1.9c-0.5,0.1-0.9,0.2-1.5,0.1    c2.6-2.1,5.4-3.3,8.4-4.3c7.2-2.5,14.6-3.9,22.1-5c12.9-2,25.9-3.1,38.9-3.5c11.5-0.4,23-0.4,34.5,0.7c5,0.5,10,1.1,14.8,2.8    c0.9,0.3,1.9,0.7,2.8,1.3c-1.4,0.3-2.5,0-3.7-0.3c-1.3-0.2-2.6-0.6-3.9-0.4c-1,0.2-2.3,0.6-2.3,1.4c0,0.9,1.3,1,2.2,1.2    c1.4,0.3,2.9,0.3,4.3,0.3c0.1,0,0.2-0.1,0.4-0.1c0.7,0.1,1.9-0.4,1.9,0.9c0.1,1.4-1.1,1-1.8,1c-8.3-0.1-16.6-0.7-24.8-1.8    c-8-1.1-15.9-1.3-23.9-0.7c-6.6,0.5-13.3,0.7-20,1c-13,0.6-25.7,2.9-38,7c-2,0.7-4.2,0.4-6.3,0.7c1.4-1.3,3.1-2.2,4.8-2.9    c3.3-1.3,6.9-1.8,10.3-2.8c1.1-0.3,2.2-0.7,3.3-1C97.4,255.2,97.1,255.2,96.9,255.1z"
	pathData6 := "M151.2,114.7c0.1,1.7,0.1,3.4-0.1,5.1c-0.5,3.1-2,5.7-4.4,7.7c-2.5,2.1-2.5,4.5,0,6.6    c0.5,0.4,1.1,0.8,1.5,1.3c1,1.1,1,2.2,0,3.2c-2.6,2.3-5.4,4.4-8.8,5.3c-1.4,0.4-2,0-2.2-1.5c-0.1-0.9-0.2-1.7-0.2-2.6    c-0.1-2.3-1.1-3.5-3.3-3.8c-1.3-0.2-2.4-0.7-3.4-1.7c-1.3-1.5-1.4-3,0.2-4.1c2-1.4,3.9-3.2,6.5-3.5c1.9-0.2,2.3-1.5,1.4-3.2    c-0.4-0.7-1-1.4-1.1-2.2c-0.2-1.4,0.5-2.1,1.8-1.8c0.5,0.1,0.9,0.2,1.4,0.3c1.8,0.2,2.6-0.5,2.6-2.3c0-1-0.1-2.1-0.1-3.1    c0-0.6,0-1.3,0-1.9c1.1,0.1,1.9,0.7,2.3,1.6c0.3,0.7,0.6,1.4,0.9,2.1c0.3,0.6,0.6,1.3,1.1,1.7c0.4,0.4,0.9,0.9,1.4,0.6    c0.5-0.2,0.2-0.9,0.2-1.3c0-2.2-0.6-4.3-0.5-6.4c0.1-1.5,0.3-2.9,1.3-4.3c0.5,1.3,1.5,2.4,1.3,3.9c-0.1,0.8-0.1,1.5,1,1.5    c1.1-0.1,1.9-0.9,1.8-1.8c-0.1-0.3-0.2-0.6-0.4-0.9c-0.9-1.7-0.9-2.6,0.5-4.4c1.5,5.9,1.7,11.8,2.3,17.9c1-1.1,1.3-2.3,1.5-3.6    c0.2-1.4,0.4-2.8,1-4.2c0.3-0.6,0.5-0.7,1-0.1c0.9,1.2,1.3,2.6,1.4,4.1c0.1,1.8-0.2,3.6,0.3,5.4c1.4-0.3,2.4-1.2,3.5-2    c1.5-1,3-1.4,4.8-1c1.6,0.3,2,1.3,1.1,2.6c-2,3.4-1.9,3.9,1.4,5.9c0.7,0.4,1.4,0.8,2.1,1.1c1,0.4,1.7,1,1.8,2.2    c0.1,1.2-0.3,2.1-1.3,2.8c-1.3,0.8-2.6,1.5-4,2.1c-2.3,0.9-2.8,1.8-2.7,4.3c0,0.9-0.3,1.1-1.1,1.2c-1.1,0.1-2.1-0.3-3.1-0.8    c-2.1-1.1-4-2.6-5.7-4.4c-0.4-0.4-0.8-0.9-1-1.5c-0.4-0.9-0.2-1.5,0.6-2.1c1.2-1,2.6-1.5,4.2-1.4c-2.3-0.4-3.4-2.2-4.5-3.9    c-0.9-1.4-1.9-2.8-3.3-3.7c-1.2-0.8-1.4-2-1.5-3.4c-0.1-1.8,0.4-3.6,0.1-5.4c-0.2-1.1-0.6-2.1-1.3-3    C151,113.8,151,113.9,151.2,114.7z"
	pathData7 := "M84.1,115.9c2.8,2.3,3.2,5.9,5.2,8.4c0.6-0.7,0.7-1.3,0.4-2.4c-0.2-0.8-0.5-1.6-0.7-2.4    c-0.2-0.8-0.2-1.6,0.4-2.5c1.4,3.2,4,5.5,4.6,8.9c0.2,1.4-0.1,2.7-0.2,4c-0.1,0.8-0.7,1-1.4,1c-1.7,0-2.6,1-3.1,2.5    c-0.2,0.6-0.4,1.2-0.7,1.8c-1.4,3.6-3.4,6.7-7.8,6.9c0.8,1.8,1.4,2.1,3.1,1.3c2.5-1.2,4.9-1.1,7.2,0.6c0.9,0.7,1.4,1.8,0.8,2.7    c-1.7,2.9-3.4,5.9-7.1,6.8c-0.4,0.1-0.9,0.2-1.3,0.2c-1.4,0-2-0.6-1.7-2c0.1-0.7,0.3-1.5,0.2-2.3c-0.1-1.4-0.8-2-2.2-1.7    c-1.3,0.3-2.5,0.1-3.8-0.1c-2-0.3-2.8-2.3-1.6-3.9c0.7-1,1.7-1.7,2.6-2.5c1-0.9,1.4-1.8,0.6-3.1c-0.1-0.2-0.2-0.5-0.3-0.8    c-0.7-2-0.2-3,1.9-3.3c1.5-0.2,2.9-0.5,4.4-0.3c0.9,0.1,1.7,0.1,2.1-0.9c0.4-1,0.2-1.8-0.6-2.5c-1.1-1.1-2.3-2-3.3-3.1    c-1.4-1.4-2.3-3-1.9-5c0.2-1.1,0.5-1.3,1.5-0.8c0.6,0.3,1,0.7,1.4,1.2c0.5,0.6,1,1.3,1.5,1.9c0.3,0.3,0.7,0.9,1.2,0.7    c0.4-0.2,0.3-0.8,0.2-1.3c0-1.2-0.3-2.4-0.7-3.5C84.6,119.1,84.2,117.7,84.1,115.9z"
	pathData8 := "M198.9,137.1c0.8,0.3,1.4,0.5,2,0.7c3.1,1,5.2,3.2,6.5,6.1c0.2,0.4,0.4,0.7-0.2,0.9c-2.2,0.7-4.3,1.8-6.6,2.2    c-1.3,0.2-2.6,0.2-3.9-0.1c-0.6-0.2-1-0.4-0.6-1.1c0.4-0.8,0.6-1.7,0.7-2.6c0.2-1.5-0.6-2.5-1.8-3.2c-0.2-0.1-0.4-0.2-0.7-0.3    c-1.3-0.5-2.5-1.3-2.5-2.8c0-1.6,1-2.6,2.3-3.4c1.2-0.7,2.6-0.7,3.9-0.9c2.8-0.2,3.6-1.3,3-4c-0.1-0.4-0.2-0.8-0.2-1.2    c0-1.1,0.6-1.7,1.6-1.4c2.7,0.7,5,2,6.5,4.5c0.7,1.2,0.5,2.5-0.2,3.7c-1.3,2.4-4.1,3.6-6.8,3.1C201,136.9,200.1,136.9,198.9,137.1    z"
	pathData9 := "M104.7,125.9c0,0.6,0,1.1,0,1.6c0.1,2.6,1.6,4.1,4.2,4.4c0.5,0,1,0,1.6,0c1.1,0.1,2,0.5,2.3,1.6    c0.3,1.1,0.3,2.1-0.6,2.8c-0.9,0.7-1.8,1.4-2.7,2.2c-2.1,1.7-2.2,3.1-0.4,5.2c0.2,0.2,0.5,0.5,0.6,0.7c0.9,1.2,0.4,2.4-1.1,2.6    c-3,0.5-5.8-0.3-8.6-1.5c-0.8-0.3-1.2-0.9-1.1-1.8c0.3-1.4,0.5-2.7,1-4c0.7-1.8,2.1-2.6,4-2.7c0.7,0,1.4-0.1,2-0.7    c0.3-0.3,0.5-0.6,0.4-1c-0.1-0.5-0.6-0.3-0.9-0.3c-1.5,0-3,0-4.4,0c-1.6,0-2.9-0.5-3.5-2.2c-0.6-1.6-0.2-3,1.3-4.3    C100.4,127.2,102.5,126.5,104.7,125.9z"
	pathData10 := "M177.5,258.6c0.6,5.4,0.7,10.6,2,15.6c0.8-1.2,1.4-3.5,1.7-6.2c0.2-2.3,0.5-4.6,0.7-6.9    c0.1-1,0.8-1.1,1.6-1.1c0.9,0,1.5,0.3,1.5,1.3c-0.1,2,0.1,3.9-0.1,5.9c-0.2,1.9,0,3.7,0.5,5.6c0.3,1.2,0,1.7-1.2,1.8    c-1.7,0.2-3.3,0.4-5,0.5c-0.7,0.1-1.4,0.1-2,0.1c-2.5-0.1-2.7-0.3-2.5-2.7c0.3-2.8,0.2-5.6,0.8-8.3    C175.8,262.3,176.4,260.5,177.5,258.6z"
	pathData11 := "M228.5,260.2c-4.4-0.1-8.2-1.9-12.1-3.4c-1.8-0.7-3.7-1.3-5.6-1.6c-0.3,0-0.6-0.3-0.8,0    c-0.1,0.2,0,0.5,0.1,0.8c0.2,0.5,0.5,0.8,1,1c0.3,0.2,0.7,0.4,1.1,0.5c0.6,0.3,1.3,0.6,1.2,1.3c-0.1,0.4-0.9,0.2-1.4,0.2    c-5.3,0-10.5-0.7-15.3-3.1c-0.2-0.1-0.5-0.3-0.7-0.4c2-0.9,4.8-0.7,11.4,0.6c-0.6-1.4-1.9-2.4-3.7-2.8c-1.3-0.2-2.7-0.3-4-0.5    c-1.6-0.2-2.5-0.8-3.4-1.9C203,250.3,223.3,256.1,228.5,260.2z"
	pathData12 := "M192.7,260.1c0.5,0,0.9,0.1,1.3,0.1c3.2,0.5,3.9,1.7,3.1,4.9c-0.4,1.6-0.4,3.3,0.3,4.9c1.1,2.4,0.5,3.4-2.1,4    c-2.5,0.6-4.9-0.1-7.4-0.1c-0.6,0-1-0.5-0.9-1.2c0.4-3.8,0.9-7.5,1.3-11.3c0.1-0.8,0.7-1.1,1.4-1.3c0.5-0.1,0.9,0,1.2,0.4    c0.3,0.6,0.3,1.1-0.2,1.7c-0.8,1-0.8,2.2-0.4,3.4c0.6,1.6,1,3.2,0.5,4.9c-0.1,0.5,0,1,0.5,1.2c0.5,0.3,1,0.2,1.5-0.1    c0.7-0.4,1-1.1,1-1.8C193.7,266.6,194.3,263.3,192.7,260.1z"
	pathData13 := "M93.7,249.7c-1.4-3.5-2.3-7.2-2.9-10.9c-0.7-3.7-1.4-7.4-1.8-11.1c-0.2-1.5-0.3-2.9-0.8-4.4    c-0.4-1.1,0.2-1.9,1.4-2.2c1.1-0.2,1.7,0.2,2,1.4c0.4,1.4,0.2,2.9,0.3,4.3c0.5,5.6,1.8,11,2.6,16.5c0.2,1.1,0.4,2.1,0.5,3.2    C95.1,247.8,94.6,248.8,93.7,249.7z"
	pathData14 := "M204.7,262.2c-2.4,3-2.4,6.1-0.9,9.3c1.3-1.4,1.6-3.1,1.8-4.8c0.1-0.9,0.2-1.7,0.3-2.6c0-0.5,0.2-1,0.8-1.2    c0.7-0.2,1.4-0.1,2,0.4c0.5,0.4,0.4,0.9,0.2,1.4c-0.1,0.2-0.2,0.4-0.3,0.7c-1,1.9-1,3.9-0.2,5.9c0.2,0.5,0.4,0.9,0.6,1.5    c-2.5,0.2-4.9,0.5-7.3,0.1c-1.3-0.2-1.8-0.8-1.7-2.1c0-0.7,0.1-1.4,0.3-2c0.4-1.9,0.8-3.7,0.2-5.6c-0.4-1.4,0.3-2,1.8-1.6    C203.1,261.6,203.8,261.9,204.7,262.2z"
	pathData15 := "M201.2,187.7c0.1,1-0.3,1.9-0.7,2.8c-1.1,2.9-2.4,5.8-2.8,9c-0.1,0.5-0.1,1-0.2,1.5c-0.2,0.8-0.2,1.6-1.3,1.8    c-1,0.1-2-0.6-2.3-1.5c-0.2-0.5-0.2-1.1,0-1.7c1.1-3.8,1.9-7.7,3.4-11.4c0.5-1.2,1.1-2.3,1.7-3.4c0.3-0.7,0.7-0.6,1.1-0.1    C200.8,185.5,201.2,186.5,201.2,187.7z"
	pathData16 := "M164,257.2c1.4,2.7,1.9,5.5,2,8.4c0.1,2.7-0.3,5.4,0.1,8.1c0.1,1-0.5,1.3-1.4,1.4c-1.4,0.1-2.4-0.5-2.6-1.7    c-0.1-0.7-0.1-1.5-0.1-2.3C161.8,266.4,162.3,261.8,164,257.2z"
	pathData17 := "M174.4,274.5c-2.5,0.9-4.9,0.4-7.5,0.4c1.8-2.9,1.6-6.1,1.8-9.2c0.1-2.6,0.6-5.1,2.4-7.4    c0.8,1.8,0.8,3.5,0.8,5.3c0,1.7-0.2,3.4-0.1,5.1C171.9,271,172.6,272.9,174.4,274.5z"
	pathData18 := "M146,263c0,2.6,0,5.2,0,7.9c0,0.6,0.1,1.3,0,1.9c-0.1,1.1-0.8,1.7-1.9,1.7c-1.2,0-1.4-0.9-1.5-1.8    c0-0.7-0.1-1.5,0-2.2c0.5-2.8,0.3-5.5-0.2-8.2c-0.1-0.8-0.1-1.6,0.2-2.4c0.3-0.9,1.1-1.3,2.1-1.2c1,0.1,1.2,0.8,1.2,1.6    C145.9,261.2,145.9,262.1,146,263C145.9,263,146,263,146,263z"
	pathData19 := "M155.8,269.7c0.1-2.7,0.2-5.4,0-8.1c-0.1-1.5,0.6-2.2,1.9-1.9c1.1,0.2,1.3,1,1.4,1.9c0.1,1.2-0.1,2.5-0.2,3.7    c-0.1,1.3,0.1,2.4,0.7,3.6c0.7,1.4,0.2,3.9-1,4.9c-0.6,0.5-1.3,0.4-1.9,0.2c-0.6-0.2-0.8-0.8-0.8-1.4    C155.8,271.6,155.8,270.6,155.8,269.7z"
	pathData20 := "M183.6,192.2c0-0.1,0-0.2,0-0.2c1,0.2,2-0.1,3,0c1.5,0.1,1.8,0.5,1.8,1.9c0,1.4,0.1,2.9-0.2,4.3    c-0.4,1.6-1.2,2.7-2.7,3.5c-1.4,0.7-2.2,0.4-2.5-1.2c-0.4-2,0.2-4,0.4-6C183.6,193.6,184.2,192.9,183.6,192.2z"
	pathData21 := "M129.8,258.1c1.5,2.2,2.1,4.3,2.5,6.5c0.5,2.5,0.9,5.1,0.6,7.7c-0.1,0.5-0.1,1-0.4,1.4    c-0.2,0.5-0.4,1.1-1.1,1.1c-0.7,0-1.1-0.5-1.4-1.1c-0.1-0.3-0.2-0.7-0.2-1C129.8,267.8,129.8,263.1,129.8,258.1z"
	pathData22 := "M95.5,137.6c-0.1,2.5-1.5,4.3-3.4,4.3c-1.9,0-3.4-1.9-3.3-4.2c0-1.6,2.1-3.5,3.7-3.4    C93.9,134.3,95.6,136.1,95.5,137.6z"
	pathData23 := "M200.9,240.8c0,0.9-0.4,1.8-0.7,2.6c-0.3,0.8-0.6,1.6-0.5,2.5c0,0.3,0.2,0.6-0.2,0.7c-0.3,0.1-0.6,0-0.7-0.3    c-0.3-0.4-0.5-0.8-0.7-1.2c-0.4-0.8-1-0.8-1.6-0.2c-0.1,0.1-0.2,0.2-0.3,0.3c-0.6,0.7-1.3,1.3-2.3,1c-1-0.3-1-1.3-1.2-2.2    c-0.2-1.1,0.3-1.8,1.4-2.1c1.8-0.5,3.6-1.3,5.3-2.1C200.7,239.2,200.9,239.3,200.9,240.8z"
	pathData24 := "M147.3,129.5c0.1-2.5,1.3-3.7,3.1-3.6c1.9,0.1,3.6,1.9,3.5,3.7c0,0.8-0.4,1.5-1.1,2c-0.9,0.7-1.9,1-3,1.2    c-0.8,0.2-1.4-0.3-1.8-0.9C147.4,131.2,147.3,130.3,147.3,129.5z"
	pathData25 := "M125.1,275.9c-1.6-2.4-2.9-4.6-1.7-7.4c1-2.4,0.6-4.9,0.1-7.3c-0.3-1.4,0.1-2,1.2-2.1c0.7-0.1,1.2,0.2,1.2,1    c-0.2,5.1,0.6,10.1-0.6,15.2C125.2,275.3,125.2,275.5,125.1,275.9z"
	pathData26 := "M152.2,266.9c-0.1,1.9-0.2,3.7-0.3,5.6c0,0.7,0,1.5-1,1.4c-0.9-0.1-1.6-0.5-1.7-1.5c0-0.8,0.1-1.7,0.1-2.5    c0.3-3.3,0.5-6.6,0.8-9.9c0-0.5,0-1,0.8-1c0.8,0,0.8,0.4,0.8,1C152,262.3,152.1,264.6,152.2,266.9z"
	pathData27 := "M191.3,246.2c-0.7-0.2-1.3-0.3-1.9-0.5c-1.9-0.7-2.8-1.9-2.6-3.9c0.1-1.2,0.4-2.4,0.6-3.7    c0.1-0.5,0.4-1,1-1.1c2-0.4,4.1,2,3.5,3.9c-0.1,0.2-0.2,0.4-0.3,0.7C190.7,243.7,190.7,243.8,191.3,246.2z"
	pathData28 := "M193.9,191.8c0,1.8-0.5,3.2-0.9,4.6c-0.4,1.2-0.9,2.4-0.8,3.7c0,0.8-0.1,1.5-0.9,1.7    c-0.8,0.2-1.4-0.3-1.9-0.8c-0.6-0.5-0.8-1.3-0.4-2c0.9-1.4,1-3.1,1.3-4.7c0.3-1.6,0.9-3,2.3-3.9c0.3-0.2,0.5-0.6,0.9-0.4    c0.4,0.2,0.3,0.7,0.4,1.1C193.9,191.6,193.9,191.9,193.9,191.8z"
	pathData29 := "M204.9,193.8c-0.2,1.8-0.9,3.8-1.6,5.7c-0.3,0.9-0.6,1.8-1,2.7c-0.3,0.6-0.7,1-1.4,0.7    c-0.7-0.2-1.3-0.6-1.1-1.5c0.6-2.8,1.2-5.6,2.3-8.3c0.3-0.8,1-1.3,1.9-1.1C204.8,192.2,205,192.8,204.9,193.8z"
	pathData30 := "M218.7,265.2c1.3,2.1,1.7,2.3,3.9,1.4c1-0.4,2.8,0.7,3,1.7c0.1,0.5-0.2,0.8-0.5,1c-1.6,1.1-3.2,1.9-5.2,1.9    c-2.5,0-3.4-1.5-2.2-3.7C218.1,266.9,218.6,266.2,218.7,265.2z"
	pathData31 := "M77.5,209.2c-0.4-1.3-0.6-2.3-0.9-3.3c-0.2-0.8-0.6-1.7-1.3-2.2c-0.8-0.6-1.5-1.1-1.6-2.2    c-0.1-0.5-0.6-0.6-1-0.4c-0.6,0.3-0.3,0.6-0.1,1c0.1,0.1,0.1,0.2,0.2,0.3c0.3,0.4,0.5,0.8-0.1,1.2c-0.6,0.4-0.9,0-1.2-0.4    c-0.5-0.7-0.7-1.5-0.9-2.3c-0.3-0.9-0.6-1.8-1.5-2.6c0.5,0,0.8-0.1,1.1,0c1.3,0.3,2.5,0.7,3.8,1.1c2.5,0.7,3.7,2.4,3.8,4.9    C77.8,205.8,77.7,207.3,77.5,209.2z"
	pathData32 := "M208.1,221.7c-0.4,3.2-0.3,6.3-0.1,9.5c0.1,1.4-0.4,2.6-1,3.9c-0.5,1-0.9,1.9-1.4,2.9    c-0.1,0.1-0.2,0.4-0.4,0.2c-0.3-5.5,0.2-11,2.4-16.1c0.1-0.2,0.2-0.4,0.3-0.6C208.1,221.4,208.2,221.5,208.1,221.7z"
	pathData33 := "M85.8,236.8c-1-4.3-2-8.5-2.9-12.7c-0.1-0.3-0.3-0.9,0.3-0.9c0.4,0,1-0.6,1.3,0.2c1.3,3.7,2.5,7.3,1.9,11.3    C86.3,235.4,86.2,236.1,85.8,236.8z"
	pathData34 := "M137.3,257.7c1.5,2.8,1.7,5.8,1.7,8.9c0,2.2-0.2,4.4,0,6.6c0.1,0.7-0.3,0.8-0.9,0.8c-0.5,0-0.8-0.1-0.8-0.7    C137.3,268.1,137.3,262.9,137.3,257.7z"
	pathData35 := "M190.6,232c-0.5-0.1-1.2,0.2-1.5-0.4c-0.4-0.8,0.2-1.3,0.6-1.8c1.2-1.6,2.8-2.7,4.9-2.7    c1.4,0,1.9,0.9,1.2,2.1C194.6,231.1,192.9,232,190.6,232z"
	pathData36 := "M210.5,271.7c2.2-1.8,1.5-3.8,0.8-5.8c-0.2-0.5-0.6-1.1-0.1-1.6c0.5-0.5,1.2-0.5,1.9-0.4    c0.7,0.2,0.7,0.8,0.8,1.3c0.2,1.8,0.5,3.6,0.7,5.3c0.1,1-0.4,1.5-1.4,1.4C212.4,272,211.6,271.8,210.5,271.7z"
	pathData37 := "M111.2,260.3c0.5,1.3,0.5,2.4,0.6,3.4c0.1,2,0,4,0.4,6c0.1,0.5,0.3,1.1,0.5,1.6c0.3,0.7,0.4,1.3-0.2,1.8    c-0.4,0.4-0.9,0.7-1.5,0.4c-0.8-0.3-1.1-0.9-1.1-1.7c0.1-1.8,0.3-3.5,0.5-5.2C110.7,264.7,110.9,262.7,111.2,260.3z"
	pathData38 := "M209.8,136.5c0-1.4,0.5-1.8,1.7-1.7c0.6,0.1,1.1,0.1,1.6,0.3c1.3,0.4,1.9,1.3,1.7,2.6c-0.1,1.3-0.8,2-2,2.1    c-1.5,0.2-2.4-0.4-2.8-1.7C209.9,137.6,209.9,137,209.8,136.5z"
	pathData39 := "M93.4,206.2c-1.8-2.5-2.5-5.1-2.6-7.8c0-0.4,0-0.7-0.1-1.1c-0.1-0.3-0.7-0.7-0.2-1c0.5-0.4,1.2-0.4,1.8-0.2    c0.4,0.1,0.7,0.4,0.7,0.9c-0.2,1.8,0.5,3.5,0.8,5.3C93.9,203.5,94,204.8,93.4,206.2z"
	pathData40 := "M118.2,273.7c-1.9-4.7-0.9-9.3-0.3-13.8C118.7,264.4,119.9,268.9,118.2,273.7z"
	pathData41 := "M88.7,247.9c-1.2-0.8-1.7-1.7-1.9-2.9c-0.3-1.6,0-3.2-0.9-4.7c-0.3-0.5,0-1.1,0.6-1.3c0.6-0.3,1,0.1,1.3,0.6    c1,1.5,1.2,3.1,1.2,4.8C89,245.5,88.8,246.6,88.7,247.9z"
	pathData42 := "M225.3,206.9c-0.7-1.7-0.2-3.3,0.2-4.8c0.4-1.2,0.8-2.4,0.8-3.7c0-0.8,0.6-1.1,1.3-1.3c0.5-0.1,1,0.1,1.2,0.5    c0.4,0.6,0.2,1.2-0.3,1.7c-1.1,1.2-1.4,2.7-1.8,4.2C226.5,204.7,226.2,205.8,225.3,206.9z"
	pathData43 := "M86.4,196.5c0.9,2.2,2,4.2,2.3,6.6c0.1,1,0.1,2.1-0.2,3.1c-1.7-1.2-2.5-2.8-2.5-4.8c0-1.5,0.1-3,0.2-4.5    C86,196.7,85.9,196.5,86.4,196.5z"
	pathData44 := "M215,220.5c0.4,4.7-1.8,8.8-2.6,13.5C211.2,228.9,213.5,224.8,215,220.5z"
	pathData45 := "M221,132.3c-0.1,1.5-0.5,2.8-1.3,3.9c-0.3,0.5-0.5,0.5-0.9,0c-1.7-2-1-4.9,1.4-6c0.4-0.2,0.5,0,0.6,0.3    C220.9,131.2,220.9,131.8,221,132.3z"
	pathData46 := "M232.1,203.2c0,2.1-0.8,3.8-1.6,5.5c-0.1,0.2-0.2,0.4-0.5,0.4c-0.4,0-0.5-0.3-0.7-0.6    c-0.4-1.2-0.3-2.3,0.2-3.4c0.5-1.1,1-2.2,1.5-3.4c0.1-0.1,0.2-0.2,0.2-0.4c0.2,0.1,0.5,0.3,0.5,0.4    C232,202.3,232,202.9,232.1,203.2z"
	pathData47 := "M203.7,220.8c0.1,0.9,0.2,1.7,0.3,2.6c0.1,0.9-0.2,1.6-0.9,2.2c-0.7,0.6-1.3,0.1-2-0.1    c-0.2-0.1-0.1-0.2,0-0.4c0.6-1,0.5-2.3,0.7-3.4c0.1-0.4,0.2-0.9,0.4-1.2c0.3-0.6,0.7-1.3,1.3-1.2    C204.1,219.5,203.7,220.3,203.7,220.8z"
	pathData48 := "M42,155c2,0,4.4,3.1,4,4.9c-0.1,0.5-0.5,0.7-0.9,0.8c-0.5,0.2-0.6-0.3-0.7-0.6c-0.8-1.4-1.7-2.7-2.7-4    c-0.2-0.2-0.6-0.5-0.4-0.8C41.4,155.1,41.8,155,42,155z"
	pathData49 := "M82.8,207.5c-0.6-0.7-0.7-1.4-0.9-2.1c-0.4-1.3-0.9-2.5-1.6-3.6c-0.2-0.3-0.3-0.6-0.4-1    c-0.3-0.8-0.1-1.4,0.7-1.7c0.8-0.3,1.4,0.1,1.7,0.9c0.4,1.3,0.6,2.7,0.6,4.1C82.8,205.2,82.8,206.2,82.8,207.5z"
	pathData50 := "M97.8,261.9c1.2,3,1.5,5.9,0.2,8.9C96.5,269.3,96.6,264.2,97.8,261.9z"
	pathData51 := "M210.3,194.2c-0.1,3.1-1.2,6-3.1,8.7C206.2,199.1,208.5,196.8,210.3,194.2z"
	pathData52 := "M104.1,263.8c1,2.8,3.2,5.6,0,8.5C104.1,269.4,104.1,266.6,104.1,263.8z"
	pathData53 := "M218.8,226.9c-0.2,2.4,1.1,5-1.5,7C216.3,231.2,217,229,218.8,226.9z"
	pathData54 := "M53.5,158.2c1.5,0.4,2.8,0.9,3.3,2.4c0.2,0.8,0.3,1.7-0.7,2.2c-0.7,0.3-1.8-0.6-2-1.5c-0.1-0.5-0.1-1-0.2-1.5    C53.8,159.4,53.7,158.9,53.5,158.2z"
	pathData55 := "M111.9,229c0.8,0,1.9-0.3,2,1c0.1,0.8-1.1,2-1.9,2c-1,0-2-0.9-2-1.9C110,228.7,111,228.9,111.9,229z"
	pathData56 := "M80.2,233.4c-0.4-1.1-0.8-2.3-1.1-3.4c-0.3-1.2-0.1-2.3,0.6-3.5c1.4,1.8,1.3,3.7,1.1,5.6    c-0.1,0.5-0.1,1-0.2,1.5C80.3,233.8,80.3,233.6,80.2,233.4z"
	pathData57 := "M58.7,152.2c-0.4,0.9-0.9,1.7-1.5,2.2c-0.5,0.4-1.2,0.7-1.8,0.3c-0.5-0.4-0.6-1.1-0.4-1.8    c0.2-0.6,0.6-1.2,1.4-1.1C57,151.9,57.8,152,58.7,152.2z"
	pathData58 := "M254.8,153.3c0.6-1.1,1.1-2,1.9-2.7c0.3-0.3,0.6-0.5,1-0.6c0.7-0.2,1.7-0.4,2.1,0.3c0.3,0.8-0.7,1.1-1.3,1.3    C257.2,152.1,256.2,153,254.8,153.3z"
	pathData59 := "M81.4,243.8c2.2,0,3.8,1.5,3.5,3.2c-0.2,0.8-0.5,0.8-1.1,0.4c-0.4-0.3-0.7-0.8-0.9-1.2    C82.5,245.3,82,244.5,81.4,243.8z"
	pathData60 := "M215.5,126c0.5,1.5,0.9,2.8-0.1,4c-0.3,0.3-0.6,0.8-1.1,0.6c-0.4-0.2-0.4-0.7-0.5-1.1c0-0.8,0.3-1.4,0.7-2    C214.8,127.1,215.1,126.7,215.5,126z"
	pathData61 := "M216.7,242.3c0.8,2.2,0,5.2-1.8,6c0.6-2.1,1.1-4,1.7-6C216.6,242.3,216.6,242.3,216.7,242.3z"
	pathData62 := "M220,201.1c-0.6,1.9-1.2,3.7-1.8,5.5C217.3,203,218,201.8,220,201.1z"
	pathData63 := "M233.5,217.9c-2.1-0.4-3-2-4.2-3.6c1.3,0,2,0.7,2.7,1.4c0.7,0.7,1.3,1.5,2.2,1.9    C234.2,218.2,233.8,217.9,233.5,217.9z"
	pathData64 := "M248.3,160.1c-0.9-2,0.2-3.4,1.2-4.8C250.5,157.3,248.9,158.4,248.3,160.1z"
	pathData65 := "M216.7,241.1c0.1-1.5,0.2-3,0.4-4.5c1.6,1.8,1.5,3,0,4.6C216.9,241.2,216.8,241.2,216.7,241.1z"
	pathData66 := "M222.3,122.8c0.5,1,0.8,2,0.6,3.1c-0.1,0.6-0.4,1.2-1.3,0.8c-0.4-0.2-0.6-0.2-0.4-0.7    C221.6,125,222,124,222.3,122.8z"
	pathData67 := "M229.8,240.5c0.3-0.3-0.4-1.1,0.5-1.2c0.6-0.1,1.6,1.1,1.5,1.8c-0.1,0.7-0.6,0.8-1.1,0.8    C229.8,241.9,229.8,241.2,229.8,240.5z"
	pathData68 := "M75.9,230.4c0,0.5,0,1.2-0.6,1.2c-0.7-0.1-1.3-0.7-1.3-1.5c0-0.6,0.3-1,1-1C75.9,229.1,75.9,229.8,75.9,230.4    z"
	pathData69 := "M106.6,231.4c-0.4,0.4-0.8,0.5-1.2,0.5c-0.6,0-1.2,0-1.5-0.7c-0.2-0.8,0.5-1,0.9-1.4c0.3-0.3,0.4,0,0.6,0.1    C105.7,230.5,106.2,230.9,106.6,231.4z"
	pathData70 := "M220.4,221.6c0.6,1.1,0.6,1.9-0.2,2.6c-0.3,0.2-0.6,0.8-1.1,0.3c-0.4-0.4-0.3-0.9,0-1.3    C219.5,222.8,220.1,222.5,220.4,221.6z"
	pathData71 := "M45.5,152.2c0.8,0.8,1.6,1.6,2.3,2.4C45.9,155.1,45.2,154.4,45.5,152.2z"
	pathData72 := "M221,243.5c1.1,2.2,1.1,3.5,0,4.8C221,246.7,221,245.2,221,243.5z"
	pathData73 := "M222,212.1c0.3,0,0.5,0,0.7-0.2c0.3,0.1,0.6,0.1,1,0.2c0.1,0.2,0,0.3-0.2,0.3    C222.9,212.3,222.4,212.4,222,212.1z"
	pathData74 := "M229.6,221.9c0.1-0.1,0.2-0.2,0.3-0.2c0.4,0.1,0.6,0.4,0.9,0.8C230.3,222.5,229.9,222.3,229.6,221.9z"
	pathData75 := "M223.4,212.3c0.1-0.1,0.2-0.2,0.2-0.3c0.4-0.1,0.6,0.2,1,0.4C224.2,212.6,223.8,212.6,223.4,212.3z"
	pathData76 := "M151.2,114.7c-0.1-0.3-0.3-0.7-0.2-1c0.1-0.4,0.3-0.3,0.4,0C151,113.9,151.4,114.4,151.2,114.7z"
	pathData77 := "M86.4,196.5c-0.2,0.1-0.1,0.3-0.2,0.5c-0.2-0.4-0.3-0.8-0.2-1.3C86.2,195.9,86.3,196.2,86.4,196.5z"
	pathData78 := "M216.7,241.1c0.1,0,0.2,0,0.3,0c-0.3,0.3-0.3,0.8-0.3,1.2l-0.1,0l-0.1,0    C216.6,241.9,216.4,241.5,216.7,241.1z"
	pathData79 := "M233.5,217.9c0.2-0.2,0.6,0,0.8-0.3c0.2,0,0.4,0.1,0.4,0.3c0,0.2-0.2,0.1-0.4,0.1    C234,218,233.7,218.1,233.5,217.9z"
	pathData80 := "M205.4,239c0,0.4,0,0.8,0,1.3c-0.1,0-0.1,0-0.2,0c0-0.4,0-0.9,0-1.3C205.3,238.9,205.4,238.8,205.4,239z"
	pathData81 := "M208.1,221.7c-0.1-0.1-0.1-0.2-0.2-0.3c0.1-0.2,0.2-0.4,0.2-0.7C208.5,221.2,208.3,221.4,208.1,221.7z"
	pathData82 := "M96.9,255.1c0.2,0,0.6-0.2,0.6,0.2c0,0.2-0.2,0.3-0.4,0.3C97.1,255.4,97,255.2,96.9,255.1z"
	pathData83 := "M205.4,239L205.4,239c-0.1,0-0.2,0-0.2,0c0-0.2,0-0.5,0-0.7c0.1-0.1,0.3-0.2,0.4-0.2    C205.6,238.3,205.5,238.6,205.4,239z"
	pathData84 := "M80.2,233.4c0.1,0.1,0.3,0.2,0.4,0.2c0,0.2,0,0.4,0,0.6C80.2,234.1,80.2,233.7,80.2,233.4z"
	pathData85 := "M137.3,229.6c2.7-1.2,5.1-2.2,7.6-3.2c2.8-1.2,5.5-2.7,7.8-4.8c1.1-1,2-1,3.2-0.1c2.7,2.1,5.8,3.3,8.9,4.5    c1.7,0.7,3.3,1.3,4.9,2.2c0.7,0.4,1.4,0.9,1.3,1.8c-0.1,0.9-0.8,1.3-1.5,1.6c-2.5,1-5.1,1.8-7.7,2.7c-1.5,0.5-3,1.1-4.4,1.9    c-2.2,1.4-4.3,1.2-6.5,0.2C146.4,234.5,142,232,137.3,229.6z"
	pathData86 := "M175.5,218.8c0.9,5.8-0.2,11.5-1.1,17.3C173.5,231.4,174,222.5,175.5,218.8z"
	pathData87 := "M181.7,232.3c-0.4,3.6-0.9,7.1-1.3,10.7c-0.6,0-1.1-0.2-1.6-0.5c1.1-1.8,1.2-3.7,1.2-5.7    c0-1.8,0.2-3.6,1.4-5.1C181.9,231.7,181.8,232,181.7,232.3z"
	pathData88 := "M182.5,221.1c0,2.5,0,5,0,7.6C180.4,227.2,181,224.1,182.5,221.1z"
	pathData89 := "M111.2,219.2c0.7,1.1,1.6,1.9,1.6,3.2c0,0.7-0.1,1.4-1,1.5c-0.9,0.1-0.8-0.7-1-1.2    C110.6,221.6,110.8,220.4,111.2,219.2z"
	pathData90 := "M173.9,239.8c0.6,0.8,0.2,1.3,0,1.9C173.6,241.1,173.8,240.6,173.9,239.8z"
	pathData91 := "M181.7,232.3c-0.1-0.2,0-0.4-0.3-0.5c0-0.2,0.2-0.4,0.4-0.3C182,231.7,181.9,232,181.7,232.3z"
	pathData92 := "M93.4,212c1.7-1.2,3.5-1.1,5.7-0.8C97.2,212.4,95.2,211.9,93.4,212z"
	pathData93 := "M99.7,254.7c1-0.6,1.7-1,2.5-1.4c0.4-0.2,1-0.4,1.2-0.1c0.3,0.4-0.1,0.9-0.4,1.2c-0.6,0.6-1.4,0.5-2.1,0.4    C100.5,254.9,100.2,254.8,99.7,254.7z"
	pathData94 := "M108.6,252.3c-0.8,1.7-1.3,1.9-3.8,1.3C106.1,253,107.3,252.8,108.6,252.3z"
	pathData95 := "M137.4,131.6c0.9-0.5,1.9-1,2.8-1.4c0.5-0.2,1-0.3,1.5,0.1c0.5,0.3,1,0.8,0.9,1.3c-0.1,0.5-0.8,0.2-1.2,0.2    c-1.3,0-2.5,0-3.8,0C137.5,131.7,137.5,131.7,137.4,131.6z"
	pathData96 := "M164.1,132.7c-0.5,0.2-0.7,0.4-1.1,0.2C163.3,132.5,163.6,132.7,164.1,132.7z"
	pathData97 := "M149.9,229.1c-1.1,0-2.2-0.4-3.4-0.2c-0.3,0-0.5-0.1-0.5-0.3c-0.1-0.3-0.2-0.7,0.2-0.9    c1.2-0.7,2.2-1.7,3.5-2.3c1.3-0.6,2.5-0.4,3.6,0.4c0.9,0.6,0.9,1.6,0.1,2.3C152.4,229,151.2,229.2,149.9,229.1z"
	pathData98 := "M162.2,230.1c-1-0.1-1.8-0.7-2.3-1.6c-0.4-0.8,0.1-1.3,0.8-1.5c0.7-0.2,3.4,0.8,3.8,1.3    c0.1,0.2,0.1,0.3,0,0.4C163.9,229.4,163.3,230,162.2,230.1z"
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+lineColor)
	canvas.Path(pathData5, "fill:"+hexcode)
	canvas.Path(pathData6, "fill:"+hexcode)
	canvas.Path(pathData7, "fill:"+hexcode)
	canvas.Path(pathData8, "fill:"+hexcode)
	canvas.Path(pathData9, "fill:"+hexcode)
	canvas.Path(pathData10, "fill:"+hexcode)
	canvas.Path(pathData11, "fill:"+hexcode)
	canvas.Path(pathData12, "fill:"+hexcode)
	canvas.Path(pathData13, "fill:"+hexcode)
	canvas.Path(pathData14, "fill:"+hexcode)
	canvas.Path(pathData15, "fill:"+hexcode)
	canvas.Path(pathData16, "fill:"+hexcode)
	canvas.Path(pathData17, "fill:"+hexcode)
	canvas.Path(pathData18, "fill:"+hexcode)
	canvas.Path(pathData19, "fill:"+hexcode)
	canvas.Path(pathData20, "fill:"+hexcode)
	canvas.Path(pathData21, "fill:"+hexcode)
	canvas.Path(pathData22, "fill:"+hexcode)
	canvas.Path(pathData23, "fill:"+hexcode)
	canvas.Path(pathData24, "fill:"+hexcode)
	canvas.Path(pathData25, "fill:"+lineColor)
	canvas.Path(pathData26, "fill:"+lineColor)
	canvas.Path(pathData27, "fill:"+hexcode)
	canvas.Path(pathData28, "fill:"+lineColor)
	canvas.Path(pathData29, "fill:"+lineColor)
	canvas.Path(pathData30, "fill:"+lineColor)
	canvas.Path(pathData31, "fill:"+lineColor)
	canvas.Path(pathData32, "fill:"+lineColor)
	canvas.Path(pathData33, "fill:"+lineColor)
	canvas.Path(pathData34, "fill:"+lineColor)
	canvas.Path(pathData35, "fill:"+lineColor)
	canvas.Path(pathData36, "fill:"+lineColor)
	canvas.Path(pathData37, "fill:"+lineColor)
	canvas.Path(pathData38, "fill:"+lineColor)
	canvas.Path(pathData39, "fill:"+lineColor)
	canvas.Path(pathData40, "fill:"+lineColor)
	canvas.Path(pathData41, "fill:"+lineColor)
	canvas.Path(pathData42, "fill:"+lineColor)
	canvas.Path(pathData43, "fill:"+lineColor)
	canvas.Path(pathData44, "fill:"+lineColor)
	canvas.Path(pathData45, "fill:"+lineColor)
	canvas.Path(pathData46, "fill:"+lineColor)
	canvas.Path(pathData47, "fill:"+lineColor)
	canvas.Path(pathData48, "fill:"+lineColor)
	canvas.Path(pathData49, "fill:"+lineColor)
	canvas.Path(pathData50, "fill:"+lineColor)
	canvas.Path(pathData51, "fill:"+lineColor)
	canvas.Path(pathData52, "fill:"+lineColor)
	canvas.Path(pathData53, "fill:"+lineColor)
	canvas.Path(pathData54, "fill:"+lineColor)
	canvas.Path(pathData55, "fill:"+lineColor)
	canvas.Path(pathData56, "fill:"+lineColor)
	canvas.Path(pathData57, "fill:"+lineColor)
	canvas.Path(pathData58, "fill:"+lineColor)
	canvas.Path(pathData59, "fill:"+lineColor)
	canvas.Path(pathData60, "fill:"+lineColor)
	canvas.Path(pathData61, "fill:"+lineColor)
	canvas.Path(pathData62, "fill:"+lineColor)
	canvas.Path(pathData63, "fill:"+lineColor)
	canvas.Path(pathData64, "fill:"+lineColor)
	canvas.Path(pathData65, "fill:"+lineColor)
	canvas.Path(pathData66, "fill:"+lineColor)
	canvas.Path(pathData67, "fill:"+lineColor)
	canvas.Path(pathData68, "fill:"+lineColor)
	canvas.Path(pathData69, "fill:"+lineColor)
	canvas.Path(pathData70, "fill:"+lineColor)
	canvas.Path(pathData71, "fill:"+lineColor)
	canvas.Path(pathData72, "fill:"+lineColor)
	canvas.Path(pathData73, "fill:"+lineColor)
	canvas.Path(pathData74, "fill:"+lineColor)
	canvas.Path(pathData75, "fill:"+lineColor)
	canvas.Path(pathData76, "fill:"+lineColor)
	canvas.Path(pathData77, "fill:"+lineColor)
	canvas.Path(pathData78, "fill:"+lineColor)
	canvas.Path(pathData79, "fill:"+lineColor)
	canvas.Path(pathData80, "fill:"+lineColor)
	canvas.Path(pathData81, "fill:"+lineColor)
	canvas.Path(pathData82, "fill:"+lineColor)
	canvas.Path(pathData83, "fill:"+lineColor)
	canvas.Path(pathData84, "fill:"+lineColor)
	canvas.Path(pathData85, "fill:"+lineColor)
	canvas.Path(pathData86, "fill:"+lineColor)
	canvas.Path(pathData87, "fill:"+lineColor)
	canvas.Path(pathData88, "fill:"+lineColor)
	canvas.Path(pathData89, "fill:"+lineColor)
	canvas.Path(pathData90, "fill:"+lineColor)
	canvas.Path(pathData91, "fill:"+lineColor)
	canvas.Path(pathData92, "fill:"+lineColor)
	canvas.Path(pathData93, "fill:"+lineColor)
	canvas.Path(pathData94, "fill:"+lineColor)
	canvas.Path(pathData95, "fill:"+lineColor)
	canvas.Path(pathData96, "fill:"+lineColor)
	canvas.Path(pathData97, "fill:"+lineColor)
	canvas.Path(pathData98, "fill:"+lineColor)

}