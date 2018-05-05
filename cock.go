package heraldry

import svg "github.com/ajstarks/svgo"

func renderCockToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M147.3,267.3c-2.7-1.3-4.5-2.9-7.7,0.8c-2.3,2.6-8.2,1.4-12.4,0.4c-4.4-1.1-7.8-0.5-11.2,2.8c-3,3-6.6,5.2-9.3,8.8    c-1.9,2.6-6.1,3.7-8,7.3c-0.8,1.5-1.7,0-1.8-1.3c-0.3-5.2-0.7-10.4,6.1-12.8c-7,0.8-13.9,0.4-18.5,6.5c-0.5,0.7-1.2,1.5-2,1.1    c-1.1-0.5-1.6-1.5-1.6-2.8c-0.1-3,6-9.7,9.4-10.1c1.1-0.1,2.3,0.6,3.4-1c-3,0-5.8,0-8.7,0c2.1-4.8,8.5-7.5,13-5.3    c5.1,2.4,10.4,1.1,15.6,1.2c2.5,0,4.2-2.7,5.4-4.7c4.3-7,8.3-14.1,11.2-21.8c0.7-2,0.6-3.1-0.8-4.3c-5.3-4.8-8.7-10.9-11.6-17.2    c-1-2.1-2.5-3.1-4.7-4.3c-13.8-7.7-24-18.9-30.3-33.5c-1.9-4.4-3.4-9-5.1-13.5c-3.9-10.5-2.2-20.9,0.3-31.3    c0.7-2.8,0.8-5.8,1.8-8.5c1.6-4.5,3.8-8.3-4.3-8c-4.3,0.1-5.4-7.6-2-12c1.2-1.5,2.7-2.5,3.4-4.7c1.7-6,1.3-6.8-4.9-7    c-1.5,0-2.8-0.4-4.2-0.9c-1.9-0.8-5.2-0.6-5.4-2.3c-0.3-2.1,2.7-3.4,4.3-5c0.9-0.9,2.2-1.2,3.4-1.8c3.5-1.6,7.4-2.2,7.5-7.9    c0-3,3.1-5.7,7.2-5.1c2.2,0.3,4.3,0,5.7-2.1c0.3-0.4,0.8-0.9,1.2-0.8c5.7,1,11.6-1.4,17.1,0.9c1.6,0.7,3.1,1,4.7,1    c1.7-0.1,3.8-0.3,4.2,1.7c0.4,2,1,4.3-1,6c-2.8,2.4-4.9,6-8.7,7.2c-2.6,0.8-1.7,2.5-1,3.6c8.9,13.9,13.7,29.7,20.8,44.4    c3.5,7.1,8.8,12.9,15.4,16.7c9.9,5.7,18.9,12.8,28.9,19.6c-3.1-10.6-2.5-20.8-2.2-31c0.4-18.6,8.6-33.1,23.5-44.3    c13.7-10.2,29-12,44.9-8.8c8.2,1.7,12.9,8.7,16,15.9c3.4,7.8,5.3,16,4.6,24.7c-0.4,4.7-1.5,9.3,0.7,14.2c1.2,2.7,0.6,6.7-0.6,10.1    c-1.6,4.4-0.8,8.7,0.2,13.6c2.1,9.8,1.3,20.3-6.9,28.1c-2.1,2-2.2,3.7-1.5,6.3c1.7,6.6,1.7,13.3,0,19.9c-1.4,5.2-4.9,8.8-10,10.3    c-2.5,0.7-3.5,2.4-3.7,4.2c-0.8,5.6-3.7,10.3-6.3,15.1c-2.1,3.9-4.8,4.5-8.1,1.7c-1.4-1.2-2.2-1.5-3.7-0.2    c-1.6,1.5-4.3,1.7-5.3,0.1c-2.3-3.7-7.3-3.9-9.4-8c-0.4,2.6-1.4,4.8-2.6,6.9c-0.6,0.9-1.1,3.4-2.9,0.9c-0.2-0.2-0.4-0.4-0.6-0.7    c-0.1,0.5-0.3,0.9-0.4,1.4c-1.5,5.4-6.1,7.2-10.6,9.1c-2.1,0.9-2.4-0.5-2.5-2.3c-0.1-3.2,0.8-6.6-0.9-9.7    c-0.3-0.5-0.2-1.5-0.6-1.8c-5.5-3.7-9-9-12-16.1c-0.8,4.9-1.1,8.2,2.6,10.5c-2.6,2.1-6.6,3.3-3.2,7.7c0.6,0.7,0.1,1.8-1.4,1.9    c-3.3,0.4-5,2.8-5.8,5.7c-1.6,5.8-5.4,10.5-7,16.3c-1.5,5.1,0.8,4.9,4.2,4.9c2.3,0,3.9-1.6,5.8-2.3c5.7-2.2,11.8,1.6,11.3,7.1    c-0.1,1.2-0.3,2.9-1.7,2.1c-5.7-3.2-11.4-1.2-17.2-0.9c-7.4,0.4-14.2,7-14,14.3c0.1,4.7-4,6.6-6,9.8c-1,1.6-2.5-0.9-3-2.1    c-2.4-5.4-0.8-10.1,2.8-14.3c0.3-0.4,0.6-0.8,0.9-1.2c0,0-0.2-0.2-0.5-0.4c-4.6,2.4-9.6,4.3-13.8,7.3c-2.8,2.1-7,0.8-9,4.3    c-0.4,0.8-1.1-0.1-1.6-0.9c-0.9-1.7-1.5-3.2-0.2-5c2.2-2.9,5.5-3.9,8.8-4.8c1.1-0.3,2.3,0,3.1-1.3c-1.3-1.3-2.6-2.3-4.5-0.7    c-0.4,0.3-0.8,0.6-1.3,0.6c-3.3,0.2-6.1,3.5-9.7,1.9c-0.2-0.1-0.3-1.5,0-1.9c3-3.1,5.2-7.3,10.3-7.7c4.7-0.4,9.4,0.3,14.2-0.7    c6.1-1.3,9.3-4.9,11.5-10.1c1.4-4.1,3.1-8.1,4.8-12c3.6-8.4,3.5-13.3-0.5-22c-1.7,1.2-2.1,3.3-3.1,4.9c-2.2,3.4-4.8,6.5-9.2,7    c-2.9,0.3-3,1.7-2.3,4.1c1.7,5.8,1.6,5.8-4.7,6.5c-2.5,0.3-4.2,2-3.9,4.3c0.3,2.6,2.8,1.7,4.5,1.5c1.8-0.2,3.4-1.3,5.1-1.9    C143.1,257.9,147.7,261.8,147.3,267.3z"
	pathData1 := "M125.9,162.1c-0.2,0.6,0.9,2.6-1.2,1c-2.1-1.6-3.4-4.3-6.1-5.3c-0.4,0.4-0.8,0.8-1.2,1.3    c2.8,2.1,5.6,4.2,8.4,6.4c-0.4,0.4-0.7,0.8-1.1,1.2c-2.7-1.1-5-2.7-7.2-4.5c-1.4,1,0.7,1.9,0,2.5c-0.8,0.8-1.4-0.2-2-0.7    c-3.3-2.8-6.3-5.7-8.2-9.6c0.2,4.8,3.9,7.6,7.3,11.5c-7.7-2.7-9.2-9.7-13.4-15.4c-0.2,6.7,7,10.7,4.4,17.4c-3.4-3.1-6.8-6-8-11.2    c-1.9,5.5,3.3,7.2,4.1,11.1c-3.2,0.2-4.2-2.4-5.8-4.4c-1.7-2.1-0.7-5.5-3.4-7.2c-1.6,1.9,1.6,3.6,0,5.7c-2.4-2.7-2.3-6.5-4.4-9.4    c2,5.4-1.5,10.8,0.9,16.3c-2.6-1.5-2.6-1.5-3.2-19.5c-0.7,0.6-1.3,1.2-1.9,1.8c-0.9-2.7-0.9-2.7-0.3-5.6c-0.6,2.7-0.6,2.7,0.5,5.5    c-3.2,5.7-2.2,11.1,0.7,16.5c1.1,2,2.1,4.4,4.6,3.4c2.9-1.2,1.5-3.8,0.6-5.9c-0.3-0.7-1.4-1.1-0.8-2.2c3,0.5,5.5,2.1,7.3,4.5    c2.7,3.6,6.7,1.5,10,2.5c1.9,0.5,1.1-3.2,1.7-5.4c2.9,2.6,5.9,5.4,9.8,1.9c0.9-0.8,1.7-0.6,2.9,0.2c1.7,1.2,3.6,3,6,0.5    c1.3-1.4,1.8,0.6,2.5,1.2c1.4,1.4,3.2,1.7,4.4,0.5c1-1-0.4-2.2-1-3.4c-1.6-3.2-5.5-7.1,2-8.4c1.2-0.2,1.9-1.7,0.4-3.1    c-1.1-1.1-2.7-2-3.2-4.2c1.8,0.2,3.3,2,3.9,1.8c2.1-0.6,4.6,5.1,6.4-0.4c0.3-0.8,1.9,0.2,2.7,0.7c8.8,5.2,17.1,11.1,24.8,17.8    c13.1,11.2,22.9,25.2,32.9,39.1c1.1,1.5,1.6,3.4,2.6,5.6c-3.7-1.4-3.3-6.2-7.7-7c4.3,7.9,11.2,13.7,11.5,23.6    c-2.5-4.3-2.9-8.6-6.7-11.2c2.5,6.1,5.1,12.2,7.9,19.1c-7.6-6.3-8.3-16.1-15-22c2.2,6.6,7.3,11.9,7.8,19.8    c-3.1-4.8-3.2-10.6-7.9-13.8c2.9,6.2,5.3,12.5,7.4,20.4c-5.7-5.6-4.1-13.7-9.9-17.7c4.2,6,3,14.1,8.2,19.6    c-6.1-3.6-5.2-11.4-10.4-16.8c1.1,8.7,9.5,15.5,1.9,23.7c0.4-9.5-1.4-18-9.3-24.5c4.5,8.7,10.1,17.3,3.9,27.4    c1.7-9.2-2.5-16.5-7.3-24.2c-1.5,2.9,1.7,4.1,0.7,6.9c-3.1-5.5-7-10.1-6.5-16.4c0.4-5.1-2-9.7-3-14.6c3,0.3,1.1,3.8,4.1,4.4    c-0.9-4.6-2.1-8.7-6.4-10.7c-8.6-4.1-12.7-12.3-17.5-19.6c-2.1-3.2-4.9-5-7.5-7.4c1.1,2.8,4.2,4.3,4.3,8c-4.1-0.6-4-4.9-6.4-6.9    c-0.3,5.1,1.9,8.2,6.3,8.9c3.2,0.5,3.3,4.1,7.1,5.6c-8.3,1.6-11.5-5-17.3-7.5c1.9,4.4,6.3,5.1,9.2,9.5    c-10.9-4.6-20.7-8.5-28.6-15.8c-0.4,1-0.8,2.8,0.6,3.5c3.6,1.8,5.6,6.3,10.4,5.8c1.6-0.2,2.7,0.6,3.9,1.5c5.6,4.4,12,7.2,18.7,9.7    c0.9,0.2,1.6,0.7,2.7,0.9c-1.1-0.2-1.9-0.8-2.8-0.8c0.2-2.2-2-1.4-2.9-2.4c-0.5-0.6-2-0.7-1.2-2c0.6-0.9,1.2-1.1,2.5-0.5    c4.3,2,8.1,4.5,10.6,8.7c1.6,2.7,1.7,6.5,6.1,7.4c1.6,0.3,1.3,3.8,1.1,6c-0.2,2.5-2.4,1.5-3.4,1.4c-6.5-1-12.9-2.7-19.1-4.9    c-10.7-3.8-20.6-8.9-30.1-14.9c-2-1.3-3.3-3-5.5-4.7c3.5-0.8,6.1,0.1,8.7,1.5c8.7,4.7,17.2,9.8,26.1,14c6.9,3.2,14,6.5,22.3,6.1    c-1.5-3.3-4.2-1.5-6.2-2.1c-10.7-3-21-7.1-30.1-13.7c-1.2-0.8-2.6-0.8-1.6-2.9c1.5-3.2,3.5-0.8,4.3-0.1c7.6,5.6,16.7,8,25.2,11.7    c0.4,0.2,0.9-0.1,2.2-0.2c-15.8-7.5-31.9-13.1-44.4-25.2c-0.7,4.6,4.1,5.4,6.1,8.2c0.7,1,2,1.6,3.2,2.2c1,0.5,2.6,0.7,2.1,2    c-0.6,1.7-2.6,1.6-3.9,1c-3.3-1.6-6.6-3.2-8.8-6.3c-0.5-0.8-1.1-1.8-2.4-1.3c-1.4,1.7,1.6,2.1,1.1,3.8c-3.1,0-5.1-1.5-6.9-4.5    c3-0.7,5.7-1.5,7.1-4c0.2-0.3-0.5-2.5-2.2-3.2c-0.9,1.6-0.8,3.9-2.5,5.1c-2.5,1.8-2.7-2.2-4.8-2c-1,1.4,2,2.4,0.4,4    c-3.7-4.6-6.9-9.8-12-13.9c-1.1,6.4-1,6.7,3,10.3c1.2-2-2.3-2.3-1.4-4.4c2.8,0.9,3.7,3.8,5.2,5.8c6.2,8,13.4,14.7,22.5,19.2    c3.7,1.8,6.4,4,6.6,8.5c-3.1,0-2.9-3.6-5-4.6c-1.6,6.4,6.3,6.7,7.9,11.8c-6.8-3.9-11.8-8.2-12.1-16.5c-0.5,0.3-0.9,0.6-1.4,0.9    c0,1.8,0,3.5,0,5.5c-2.7-2.2-2.8-2.3-1.6-5c0.5-1.1-0.3-1.5-0.5-2c-0.7-1.5-1.2,0.1-1.6,0.4c-0.6,0.5-1.2,1.3-1.2,2    c0.1,1.9,2.2,4-0.5,5.6c-0.9-1.6-2-3.2-2.5-4.9c-0.7-2.7-2.4-2.5-5.2-1.6c2.4-2.6,4.2-3.2,6.8-1.1c-0.8-1.9-2-3.2-1.9-5.4    c-2,2.1-3.5,5.9-5.1,0c-0.9-3.2-5.4-4.7-6.5-8.5c-0.1,1.1-0.3,2.2-0.4,3.3c-2.1,0.6-1.1-3.2-4-2.2c1.7,4.2,5.2,5.7,9.1,7    c-4.4,3.3-6.1-0.8-8-2.9c-7.9-8.7-11.4-19.7-14.7-30.5c-3-9.7-1.3-19.7,1.1-29.5c0.6-2.7,1-5.4,1.8-8c1.1-3.4,1.8-7,5.8-9    c2.7-1.4,4.5-4.3,4.1-7.8c-0.2-1.4,0.2-2.2,1.7-3c6-3,6.6-5.8,3.4-12.3c-0.7-1.4-1.5-2.4-2.8-3.1c0.1,4.6,5.6,9.4-0.7,14    c-1.3-3.6-0.4-6.2,1.7-8.6c-3.4-0.1-4,2-4.3,4.3c-0.2,1.1-0.1,2.3,0,3.5c0.2,2.4-1,3.8-3,4.5c-1.4,0.5-3,0.1-3.9-1.4    c-1.3-2-1.2-4.5,0.6-5.5c1.8-0.9,0.9-5.7,4.7-3.2c0.2-0.4,0.5-0.8,0.6-1.2c0.2-1.3-0.4-1.7-1.7-1.7c-2.8,0-4.1-1.5-4.3-4.9    c3,2.4,5.8,3,7.7-0.4c1.2-2.1,0.2-4.2-1.7-5.1c-4.2-2-5.6,1.6-7.5,4.2c-0.9-0.8-1.9-1.7-3-2.7c-1.5,3.1,2.1,3.8,2.5,5.8    c0.4,2.3-1.7,3.7-2.5,6.7c-0.8-9-8.7-3.2-11.9-7.1c4.6-1.5,8.7,0.5,12.9,1.4c-3.8-5.2-9.4-1.5-14.2-3.2c2.4-2.8,6-4.1,9.2-4.4    c5.5-0.5,10.9-1.3,15.3-3.9c-1.6-1.2-1.6-1.2-5.6,0.1c-1.8,0.6-3.7,1.3-5-0.2c-1-1.1-1.8-2.9-1-4.6c0.9-1.9,2.5-3,4.8-2.7    c3.2,0.4,5.9-0.6,8.8-2.1c3.2-1.7,7.1-0.3,10.6-0.4c3,0,6.5,1.7,9.9,1.5c1.7-0.1,1.9,0.9,1.7,2.2c-0.6,3.1-6,8.1-9.1,7.3    c-3.4-0.8-5.6,1.5-9.7,1.8c11.1,6.4,14.5,17.1,18.5,27.2c2.8,7,6.6,13.5,8.4,20.9c1.4,5.8,5.4,10,9.7,13.8    c2.3,2.1,4.6,4.2,6.3,7.1c-6.2-1.4-10.6-5.5-14.8-9.9c2.4,5,3.8,10.5,8.4,14.4c-2.8,1-4.1-1.5-5.1-3c-1.2-1.8-2.5-3.2-4.7-5.3    c1.3,4.8,5.3,7.5,5.4,13.1c-3.7-5.7-5.7-11.2-9.6-15.4c-0.4,0.2-0.7,0.4-1.1,0.6c4.3,7.3,8.7,14.7,13.1,22.2    c-3.4,0.2-2.7-4.9-5.9-4.6c-1-2.8-3.5-4.8-4.7-7.4C122.3,157.5,124.4,159.6,125.9,162.1z"
	pathData2 := "M207.8,151.8c-0.5-0.1-2.2,0.5-1.2-0.7c1.2-1.7,3.2-2.7,4.4-4.5c-3.4,0.3-6,2.7-6.7,5    c-1.8,5.9-7.1,10-7.5,16.5c-0.2,2.3-2.1,4.3-4.5,5.7c0.8-6.9,1.3-13.6,4.6-19.8c2.7-5,6.9-8.2,11.1-11.5c0.7-0.6,1.8-0.8,2.3-0.2    c0.7,1-2.3,0.6-1.1,2.2c3.1,0,4.2-3.2,6.7-4.1c2-0.7,1.8-1,4.5,0.4c0.6,0.3,1.7,0.3,2.3-0.1c6.8-3.9,13.7-1.1,18.8,2.1    c4,2.6,8.2,6.1,11.6,10c0.2,0.2,0.7,0.6,0.6,0.7c-2.7,5.2,2.6,8.4,3.3,12.7c0.1,0.9,1.2,1.8-0.2,2.4c-1.3,0.4-1.4-0.8-1.9-1.6    c-2.1-3.5-2.9-8-8.1-10.5c4.8,8.3,9.4,15.7,9.2,25.9c-2.5-3.4-2.6-6.9-5.3-9.4c-2.5-2.2-4.7-5.5-3.8-9.6c-2.4,0-1.6,2.9-3.8,3.3    c0-2.7,0-5.4,0-8.8c-2.4,2.3-4.1,3.2-4.3-0.7c0-0.7-0.8-1.4-1.3-2c-0.8-1-2-1.7-3-0.9c-1.1,0.9,1.5,1.4,0.8,2.7    c-4.9-1.7-9.7-2.5-15-1.2c-6.5,1.5-9.4,7.1-13.7,10.8c-3.9,3.4-3,4.4-7.4,4.9c0.3,0.9,2.7,0.9,1.3,2.4c-3,3.2-6.3,5.8-11.4,4.3    c2,1.4,0.7,3.4,1.1,4.9c3-0.8,5.6-7,8.3-0.1c0.2,0.4,1.7,0.1,2.6-0.7c6.2-5.3,12.6-6.3,19.1-0.4c0.3,0.3,0.9,0.2,1.4,0.2    c0.3,0,0.6-0.2,0.9-0.3c-2.6-2.9-2.6-2.9-21.3-3.9c3.1-1.1,5.5-1.8,7.7-2.9c0.8-0.4,3.4-1.3,0.6-2.9c-1.3-0.8,0.1-1.3,0.5-1.7    c3.3-4,6.6-8.1,10-12c2-2.2,4.7-1.1,7-1c1.5,0.1,2.9,1.1,4.8,1.1c2.1-0.1,2.8,4,1.3,5.9c-1.5,1.8-4.3,2.1-5.5,5.4    c2.9-1.7,5.1-3.1,7.1-5c2.5-2.4,5-1.7,5.8,1.8c0.3,1,0.5,2.5,1.3,2.4c5-0.4,4.5,5,7.4,6.9c1.4,0.9-0.4,2.5-1.2,3.7    c-0.6,1-3.5,1.4-1.4,3.1c1.2,1,1.9-1.1,2.8-1.7c0.5-0.3,1-0.9,1.5-1.4c0.4,2.6-2.6,8.2-4.6,8.4c-1.3,0.1-1.7-1.3-2.5-2    c2.8-2.8,0.2-2.8-1.7-3.8c-1.4-0.7-2.8-2.1,0.2-3.3c1.6-0.6,2.8-2,5-3.7c-3.3,0.7-6.1-0.6-8.4,1.8c0.8-1.3,1.6-2.7,2.4-4    c0.5-0.8,2.3-1.6,0.5-2.6c-1-0.6-2.3-2-3.8-0.1c-1.1,1.4-2.8,2.3-4,3.4c-2,1.9-3.3,1.3-5.3,0.1c2.9-2.4,6.3-4.3,8.4-8.7    c-3.1,2-5.3,3.8-7.8,5c-0.7,0.4-2.2,1.8-3.3,0c-1-1.6,0.5-2,1.1-2.8c0.4-0.5,0.9-1,1.1-1.8c-3.4,1.3-5.4,5.1-9.7,5.3    c1.1-2.3,2.7-4.2,4.1-6.2c1.4-2.1,3.2-3.7,4.3-6.1c-6.3,3.6-9.1,11.1-15.4,15c5.1-0.4,10.1-1,14.8,1.7c4.3,2.6,10.3,2.3,13,7.7    c-1.9,1.6-3.5-1.9-5.4-0.2c3.4,2.3,7.8,3.3,10.5,6.7c1,1.3,2.8,2.5-0.7,2.4c-1.3,0-1.5,0.5,0.1,1.9c7.1,6.2,4.8,14.1,4,21.7    c-0.1,0.9-0.2,1.6-0.9,2.2c-2.4,1.9-3.2,6.9-6.2,6.1c-2.2-0.6-0.8-5.3-0.8-8.2c0-3.8,0.1-7.6-2.3-10.8c-1.2-1.7-0.8-3.4-0.5-5.9    c2.4,4,6.1,6.7,5,11.7c-0.1,0.4,0.8,1,1.2,1.4c0.5-0.9,1.2-1.6,1-2.8c-0.7-4.7-3.1-8.4-6-11.9c-7.6-9.2-17.3-13.4-29.2-13    c-3.5,0.1-7.1-2.4-10.6,0.5c-0.3,0.2-1.5-0.8-2.7-1.4c0.5,2.5,0.5,4.8-1.6,6.4c-2.3-2.2,2.4-5.2-1-7.4c-0.6,2.2-1.8,4.1-1.5,6.5    c-2.1-0.5-2-1.8-1.7-3.4c0.3-1.6-0.7-2.5-2.2-1.9c-0.6,0.2-0.9,0.2-1.1-0.1c-1-2.3,1.4-3.9,1.5-6.1c0-0.5,0.9-2.2,2.8-1.6    c3.2,1,6.3-2.4,6.4-3.6c0.5-8.4,5.9-14,10-20.3c5-1.3,8.8-5.7,12.1-5.4C216.5,146.2,212.6,150,207.8,151.8z"
	pathData3 := "M189.3,113.3c-5,3.1-5.4,9-8.1,13.5c-1,1.6-1.7,3.7-2.8,6.5c-0.8-4.5,0.5-8.4-2.3-12.2    c-0.8,10,3,19.7-0.7,28.7c-2.5-4.3,2.4-9.7-0.9-14.6c-1.6,2.4,1,5.8-2.5,7.7c1-6.8,2-13.3,2.9-19.8c0.3-2.3-0.6-5.4,3.4-5.2    c-0.5-2.6,2.3-4,2.5-6.2c0.3-2.7,1.3-5.2,3-7c12.1-12.9,26.1-21.9,44.9-20.8c8.2,0.5,12.4,2.4,19,9.5c-1.2,1.6-2.1-1.2-3.2-0.3    c-0.2,2.8,2.9,1.1,3.7,2.6c0.9,1.6,3.6,1.1,3.7,3.6c-1.1,1-2.3-1.6-3.4,0.2c1.4,1.5,4.9,2.6,2.1,6.5c1.5,2.3,5.6,4,5.3,8.3    c-0.3,3.8,2.1,7.3,0.7,11.2c-1-1.1-2-2.2-3.8-4.3c2.2,6.9,0.2,12.1-2.6,17.3c1.9-9.1-3.8-15.8-6.4-23.5c-1.4,1-0.4,2.4-1.2,3.4    c-4.1-3.6,1-10-4-13.4c0.4,2.3,2.5,4.5-0.3,6.6c-2.7-8.8-5.7-10.7-14.5-9.7c-1.8,0.2-3.8-0.7-5.7-0.9c-2-0.2-4-0.2-5.6,1.4    c0.6,1.4,2.5-0.3,3.1,1.2c-9.6,0.6-15.7,6.1-20.1,14.3c-3.2,5.9-6.9,11.6-10.2,17.5c-2,3.6-3.9,7.3-2.3,11.7c0.1,0.4,0,1-0.2,1.4    c-0.1,0.2-0.5,0.3-0.8,0.4c-1.6-4.8-1.6-4.8,3.9-18.6c-3.7,2.7-4.7,6.1-5.7,9.2c-1.4,4.1-0.9,8.4-1.2,12.7    c-0.1,1.8,2.1,5.4-2.6,5.1c0,1.1-0.5,2.5,0,3.2c1,1.7,0.8,3.3,0.4,5.2c-1.9-4.1-3.8-7.9,1.1-11.1c0.4-0.2,0.3-1.3,0.2-1.9    c-1.2-14.1,4.1-26.3,10.9-38.1c0.2-0.4,0.2-1,0.3-1.4c1.5-3.2,3-6.5,4.9-8C192.3,106.8,190.4,109.8,189.3,113.3z"
	pathData4 := "M243.3,136.5c0.4,1.3,2.9,2.3,1.2,4.3c-3.9-2.6-0.9-9.5-7.1-10.8c0,2,3.1,2.5,2.1,6.2c-1.9-2.4-3-4.1-4.4-5.7    c-1.3-1.5-3.1-2.9-5.1-2.5c-5.6,1.1-11.8,0-16.8,3.9c-0.5,0.4-1.3,0.5-1.6,1c-6.8,9.7-18.4,16.2-20.3,29.3    c-0.3,1.8-0.4,3.6-1.1,5.2c-3.1-12.7,1.8-23.3,9.9-33c-6.5,4.4-12.4,15.6-13.1,24.2c-0.2,2.8-1.3,5.6-0.1,8.5    c0.6,1.5-0.2,2.9-1.5,3.9c-1.5,1.2-2.7,2.6-2.7,4.9c-3.1-2.4-1.5-4.8-0.3-6.8c2.6-4,3.5-7.9,1.1-12.3c-0.9-1.6-1.8-3.7,1.8-3.1    c0.9,0.1,0.9-1.1,0.7-1.5c-2.3-4.6,1.2-8.7,1.2-13.1c0-3.2,0-7.5,5-8.3c0.3,0,0.6-0.6,0.8-1c4.6-8,8.6-11,17.5-12.8    c-0.3,2.1-2.2,2.7-4.1,3.7c1.6,1.3,2.7,0.3,3.2-0.1c10.2-6.8,25.2-5.6,34.3,2.7c2.7,2.5,2.7,6.1,4,9.1c0.6,1.5,0.1,3.4-1,4.6    C245.5,138.6,245.1,135,243.3,136.5z"
	pathData5 := "M238.2,215.3c-2.3-2.3-2.6-5.3-6.2-5.4c2.2,3.5,5.4,5.8,5.1,9.9c-0.1,1.4,0.4,3.4-1.4,3.9    c-1,0.3-0.7-2.4-3.1-2.2c4.9,8.4,0.3,15.2-3.5,22.3c-0.7,1.4-1.4,2.5-2.9,3c-0.2,0.1-0.9-0.6-1-1c-0.5-1.7-0.2-3,1-4.5    c3-3.7,7.2-7,6.6-13.8c-1.7,4.6-2.9,8.3-6.5,10.5c-0.7-5.8,5.5-8.9,5.7-14.8c-2.8,1.2-1.6,4.3-5.1,6.4c6.1-10.2-2.9-15.7-5.4-23.4    c-4.3-0.5-9.3-1.9-14.2-4.1c-2.1-1-4.7-2.4-4.1-4.9c0.8-3.8-0.7-3.5-3.1-2.7c-0.3,0.1-0.6,0.4-0.9,0.5c-2.1,0.8-4-0.7-4.8-1.7    c-1-1.2,1.2-2.5,2.3-3.1c0.8-0.4,1.7-1.9,3-0.9c-0.2,1.8-3,2.1-2.5,4.6c2.7,1,7-9.8,8,1.1c3-2.1,1.6-3.9-0.5-5.8    c1.2-0.1,1.8-0.1,2.4-0.2c1.9-0.3,4.2,0.3,4.6,1.8c0.5,1.8-1.8,2.9-3.2,4c-1.4,1-3.4,1.1-4.5,3.1c4.2,0.2,6.6-2.6,9.3-4.8    c1.4-1.1,1.9-1.8,4.2-0.6c4.3,2.2,8.2,5.1,12.1,7.9C234.3,203.7,237.5,208.5,238.2,215.3z"
	pathData6 := "M147.3,267.3c-1.3-6.2-5.2-8.3-10.7-5.6c-0.8,0.4-1.8,0.7-2.7,0.8c-2.7,0.5-6.4,1.9-7.6-1.1    c-1.4-3.3,2.2-5.3,4.3-7.6c1.2-1.3,1.7,0.1,2.3,0.6c1.2,0.9,2.7,1,3.6,0.2c1-0.9,0.6-2.4,0-3.7c-0.7-1.3-1.6-2.5-1.8-3.9    c-0.3-1.6,0.2-3.9,2.2-3.5c5.9,1.2,8.1-2.8,10.6-6.6c1.4-2.1,2.5-4.4,3.8-6.8c6.7,11.8,5,25.9-4.2,37    C147.2,267.2,147.3,267.3,147.3,267.3z"
	pathData7 := "M141.5,209.9c6.4,2.5,12.6,5,19.3,6.3c-1,1.3-3.2,0.4-3.6,2.3c2.4,2.1,7.5-1.4,8.8,3.9    c-1.1,0.4-1.4-0.5-2-0.9c-0.5-0.4-1-0.9-1.5-1.4c-2.6,4-2.2,5.1,2.9,8.1c1.2,0.7,0.7,2.6-0.6,3.9c-1.1,1.1,0,2.7,0.8,2.7    c2.3,0,3.1,2.2,5,2.7c-2.1,1.5-3-0.7-5.1-1.5c2.4,3.1-0.2,7,3.4,9.6c-5.3,1.1-7.4-3.5-8.4-5.9c-4.2-9-10.3-16.3-16.9-23.5    C142.2,214.6,140.1,213.2,141.5,209.9z"
	pathData8 := "M124.8,290.8c0.7-2.9,2.8-3.2,4.6-3.8c12.5-4.5,19-14.3,23.6-25.9c1.6-4,3.5-7.9,5-11.9    c0.6-1.7,1.7-1.4,2.6-0.9c1.9,1,2,2.9,1.1,4.3c-3.3,5.4-5,11.4-7.6,17c-2,4.4-0.6,6.1,4.3,6.4c2.1,0.1,4.4-0.3,5.6-1.3    c4-3.2,4.7-1.2,4.9,2.7c-5-1.7-9.8,0.5-14.6,0.9c-6.1,0.5-10.8,6.6-11.2,13.1c-0.3,4.8-0.3,4.8-6.4,3.9c3.5-5.6,7.3-10.9,9.6-18    C142.5,281.9,130.5,289.2,124.8,290.8z"
	pathData9 := "M133.4,233.8c1.7,1.9,3.1,3.4,4.4,4.9c-3.8,3.9-8.4,7.8-3.2,14.2c-1.3-1.2-2.6-2.3-3.9-3.5    c-1.9,2.9-3.9,5.6-5.6,8.6c-2.6,4.4-1,6.9,4.2,7.2c2.4,0.1,4.4-0.9,6.6-1.4c1.8-0.4,2.5,0.8,2.7,1.8c0.2,1-0.9,1.6-2,1.5    c-2.6-0.3-5.2,0.1-7.8-0.5c-4.7-1.2-9.5-0.9-12.9,3.3c-1.2,1.5-2.7,2.6-4.4,3c-3.6,1-5.2,3.5-5.9,7.3c-0.2-3.4-2.6-2.6-4.7-2.3    c2.5-4.3,7.6-4.8,10.8-9.4c-7,3.5-14,2.1-20.1,4.8c-2.5,1.1-1.9-1.3-2.2-1.8c-0.8-1.4,0.7-1.5,1.6-1.4c4.6,0.4,9.2-0.8,13.7-1.2    c7.5-0.8,13.6-4.8,17.3-11.6c3.7-6.7,8.1-13.1,10.7-20.4C132.9,236,133,235.2,133.4,233.8z"
	pathData10 := "M214.8,208.8c-0.3-1.7-1.8-1.6-3.2-3.3c5,1.6,8.4,3.3,9.2,8.6c-3.1-1.1-5.2-3.8-8.8-4.1    c12.8,11.3,12.9,11.5,12.8,18.4c-2.7-6.8-6.5-10.6-13.6-12.8c4.3,3.9,8.9,5.4,10.4,10.3c1.4,4.6,2.5,9.1,0,14.1    c-0.9-3.2,0.1-6-0.8-8.8c-1.2-3.7-3-7.2-6.3-9c-4.9-2.6-6.6-7-8.1-11.8c1.4-2,3.1,1.7,5.1-0.5c-5.8-2.8-11.3-6-12.7-12.9    C202.4,203,207.2,207.7,214.8,208.8z"
	pathData11 := "M199.3,117.8c4.4-6.8,9.9-11.5,18.4-11.8c3.6-0.2,7-0.8,8.7,3.8c0.8,2.3,1.1,2.8-1.7,3.4    c-5.6,1.4-11.3,0.5-17,1.7C204.6,115.4,202,116.5,199.3,117.8z"
	pathData12 := "M84,113.2c0.7-1.7,1.2-2.8,1.8-4.2c-2.3,0.4-3.2,1.7-3.8,3.2c-1.1,2.4-3,2.3-4.9,1.4    c-2.1-0.9-3.7-2.7-2.9-4.9c0.9-2.3,1.2-5.1,3.4-6.8c2.4-2,4.4-4.4,5.3-7.4c0.4-1.4,1.2-2.9,2.6-2.1c1.3,0.7,0.4,2.3-0.4,3.4    c-0.4,0.5-0.5,1.2-0.9,1.7c-1.8,1.9-3.1,4.7-1.9,6.7c1.3,2.3,4.6,3.3,7.2,2.6c2.4-0.6,1.8,1.1,1.5,1.5    C89.2,110.3,87.8,112.9,84,113.2z"
	pathData13 := "M139.1,219c5.6,3.2,8.1,12.1,5.4,17.7c-0.6-1.9-1.3-3-2.4-0.3c-0.6,1.4,0.7,3.3-1.4,4.9    c-0.9-3.3-1.6-6.4-4.5-8.9c3.4-0.8,3.7,1.9,5.4,2.3c2-4.7-5.5-11.9-12.7-12.1c0.6,2.6,4.4,0.1,4.7,2.9c0,0.2-1.5,0.6-2.3,0.8    c0.5,1.9,3.2,1.3,3.5,3.3c-3.2-0.8-8.9-7.1-8.7-9.6c0-0.4,0.1-0.6,0.8-0.7c1.4-0.1,1.4-0.4,0.6-2.5c3.1,6.6,11.4,4.3,15.2,10.8    C143.3,223.4,140.3,221.8,139.1,219z"
	pathData14 := "M234.7,201.7c-3.3-3.8-10.3-4.2-10.4-11.7c-1.6-0.9-1.9,2.5-4,1.7c0.5-0.8,1-1.6,1.5-2.5    c-6.2,1.5-11-4.8-17.2-2.3c8.4-2,16.1,0.6,23.8,3.2c1.9,0.6,1.4,1.7,0.2,2.9c-0.4,0.4-2.3,0.9-0.7,1.9c0.5,0.3,1.4,1.7,2.4,0.3    c0.3-0.5,0.7-1.6,1.8-0.6c0.9,0.9,0.4,1.3-0.2,1.9c-0.4,0.5-1.6,1.6,0.2,1.3C235.7,197.5,234.7,199.8,234.7,201.7z"
	pathData15 := "M255.9,137c2.5,4.4,0.7,8.7,0.1,12.9c-0.2,1.6-1.2,0.3-1.7,0.1c-4.6-1.8-5.1-4.6-1.8-8    C253.8,140.6,254.6,138.8,255.9,137z"
	pathData16 := "M217.5,136.7c1.9-4.8,5.3-6.7,10.1-5.7c1.5,0.3,2.4,1,2.4,2.5c0,1.3-0.8,2.1-2,2.4c-0.3,0.1-0.7,0.1-1,0    C223.7,135.5,220.8,138.1,217.5,136.7z"
	pathData17 := "M139.8,278.9c-3.2,1.8-6.4,3.6-9.5,5.4c-1.2,0.7-1.4,0.8-2.5-0.6c-1.2-1.7-3.6-2.8-6-1    c-0.4,0.3-1,1.8-1.6,0.5c-0.4-0.8-1.5-1.5-0.9-2.6c0.5-0.9,1.4-0.5,2.2-0.5C127.5,279.8,133.8,281,139.8,278.9z"
	pathData18 := "M109.7,264.9c-7,3.6-12.2,2.8-14.7-2.1C99.5,264.9,103.9,265.3,109.7,264.9z"
	pathData19 := "M116.6,295.8c-0.5-4.6,2.8-5.5,6.1-6.6C124.2,294.7,116.8,291.9,116.6,295.8z"
	pathData20 := "M233.5,110.4c3,1.2,2,4.8,4.2,6.3C234.2,116,229.7,116,233.5,110.4z"
	pathData21 := "M140.2,216.1c-4.2-2.3-4.8-6.4-7.6-9.6C135.9,207.6,137.5,209.5,140.2,216.1z"
	pathData22 := "M215.4,238.5c-0.2-3.4-3-5.6-3.4-8.8c-0.1-0.8,0.1-1.6,0.2-3C214.3,230.8,217,234,215.4,238.5z"
	pathData23 := "M184,179.2c0.6-3.9,2.6-6.3,5.8-8.1C189.3,174.3,188.1,175.9,184,179.2z"
	pathData24 := "M229.8,114.8c-1-3.6-1.8-6.5-2.7-9.7C230.8,107.5,232.2,110.2,229.8,114.8z"
	pathData25 := "M238.8,137.6c-2.2,0.3-4.6,1.3-5.7-1.3c-0.4-0.9,0.5-1.8,1.2-1.9C237.1,133.7,237.1,136.7,238.8,137.6z"
	pathData26 := "M217.2,133.2c-1.8,2.2-2.2,5.6-6.3,4.5C213,136.1,213.7,133.2,217.2,133.2z"
	pathData27 := "M83.6,278.7c-1.2-4.2,1.2-5.4,4.3-6.4C88,275.8,81.9,274.5,83.6,278.7z"
	pathData28 := "M138.5,301.9c-0.3-1.6-2.7-2.6-0.5-4.4c0.8-0.6,1.6-0.7,2.1,0.1C141.1,299.4,139.9,300.6,138.5,301.9z"
	pathData29 := "M88.1,266c1.1-2.1,2.8-2.5,4.5-2.8c0,0,0.6,1.3,0.3,1.7C91.8,266.9,89.9,265.6,88.1,266z"
	pathData30 := "M113.4,285.7c1.5-0.9,1.3-4.2,4.1-2.9c0.4,0.2,1.2,0.5,1.2,0.6C117.7,285.9,115,284.5,113.4,285.7z"
	pathData31 := "M175.9,278c-1.6-1.1-2.7-2.4-4.5-2c-1.1,0.2-1.4-0.3-0.7-1.3c0.4-0.6,0.4-2.3,1.6-0.7    C173.4,275.2,176,274.9,175.9,278z"
	pathData32 := "M99.1,282.3c-0.6-2.7,0.2-4.1,4-2.9C101.6,280.5,100.3,281.4,99.1,282.3z"
	pathData33 := "M142.9,263.6c-1.7,0.8-2.8,0.3-3.6-1C140.8,261.3,141.8,262,142.9,263.6z"
	pathData34 := "M246.9,139.3c1.1,1.8,2.1,2.4,1.5,3.7c-0.2,0.3-0.8,0.5-1.2,0.7C247.2,142.6,247.1,141.6,246.9,139.3z"
	pathData35 := "M87.5,85.9c-0.4-0.3-0.9-0.6-1.1-1c-0.5-1.1,0.6-1.3,1.1-1.7c0.3-0.2,1.4-0.5,1.1,0.6    C88.4,84.5,87.9,85.1,87.5,85.9z"
	pathData36 := "M180,218.5c0-2.7,0-4.4,0-6.1c-2.6,1.1-2.9,0.5-5.8-11.2c1.5,0.9,2.4,1.8,2.9,3.6c0.5,2.1,1.2,4.2,3,5.9    C182.8,213.2,182.6,213.4,180,218.5z"
	pathData37 := "M182.3,197.3c-0.1-2.4-2.8-4.1-2.1-6.9c2.2,0.5,1.9,2.7,2.9,3.8c0.8,0.9,0.4,2.6,2.3,2.8c5.2,0.5,3.1,4.1,3.8,7.2    c-2.9-0.7-1.9-2.2-2.2-3.1c-0.2-0.6-0.3-1.1-0.4-1.7c-3.4,3-2.8-1.9-4.4-2.3L182.3,197.3z"
	pathData38 := "M182.1,197.1c-0.2,1.1-0.9,3-1.8,1.3c-2-3.7-5.3-6.3-8.2-10.9c6,1.8,6.3,7.6,10.2,9.7    C182.3,197.3,182.1,197.1,182.1,197.1z"
	pathData39 := "M125.9,162.1c-3.7-2.7-5.5-6.9-7.4-10.9c2.9,3.4,6.5,6.2,7.3,11C125.7,162.3,125.9,162.1,125.9,162.1z"
	pathData40 := "M84.1,150.9c-1.9,0.4-2.1-0.7-2.1-2.2c0-2.7,0.4-5.4,2.6-7.7c1.3,3.7-1.4,6.8-0.6,10.1C83.9,151.1,84.1,150.9,84.1,150.9z    "
	pathData41 := "M136.3,184.3c-2.4-1.9-4.9-3.9-7.3-5.9c-0.5-0.4-1-0.8-0.3-1.5c0.8-0.9,1.2-0.5,1.9,0.1c2.2,2.1,4.5,4.2,6.7,6.3    C136.9,183.6,136.6,183.9,136.3,184.3z"
	pathData42 := "M83.9,133.1c0.8,0,2.3-0.2,1.9,0.7c-0.9,2.4-0.9,5.4-3.6,7.1c-1.1-3.2,2.6-5,1.8-7.9L83.9,133.1z"
	pathData43 := "M84.1,132.9c-0.6-0.3-2,1-1.4-1.2c0.6-2.3,0.6-4.8,3.1-6.9c0.6,3.3-1.8,5.6-1.9,8.2C83.9,133.1,84.1,132.9,84.1,132.9z"
	pathData44 := "M111,169.4c1.9,2.6,3.3,4.3,4.6,6.1C112.1,175.5,111.9,175.4,111,169.4z"
	pathData45 := "M176.2,182.3c-1.3-0.4-1.8-0.5-1.8-0.6c-0.2-1.7-2.4-2.3-2.2-4.1c0-0.2,0.9-0.6,1.2-0.5    C175.1,178.1,175.5,179.8,176.2,182.3z"
	pathData46 := "M151.7,195.3c2.6-1.2,4.2,0.9,6.1,2c-2.2-0.1-4.5,0-5.9-2.2C151.8,195.2,151.7,195.3,151.7,195.3z"
	pathData47 := "M194.2,214.1c-1.6-0.7-0.5-1.6-0.2-1.9c1.2-1.5-2.5-2.7-0.2-4.1c0.2-0.1,0.8,0.1,0.8,0.2    C195,210.3,195.6,212.3,194.2,214.1z"
	pathData48 := "M190.1,207.1c0-2.2,0-3.8,0-5.8C191.7,203.2,193.3,204.7,190.1,207.1z"
	pathData49 := "M185,221.8c-0.3,1.3-1,2-2,2.2c-0.9,0.2-1.2-0.7-1-1.2c0.4-0.8,1.1-1.5,1.7-2.3C184.3,220.9,184.7,221.4,185,221.8z"
	pathData50 := "M207.8,151.8c5.8-4.9,14.3-7.5,18.3-5c-6.1-0.2-11.1,1.3-15.6,4.8c-0.6,0.5-1.9,0-2.9,0    C207.7,151.7,207.8,151.8,207.8,151.8z"
	pathData51 := "M196.3,163.6c-0.5-5.7,4.1-9.1,6.4-14.8C202.7,155.7,197.9,158.9,196.3,163.6z"
	pathData52 := "M243.4,215.7c4.2-3.9,1.3-8.7,1.6-13.4C248.4,207,247.7,213.3,243.4,215.7z"
	pathData53 := "M189.3,113.3c0-5.9,4.8-8.5,8.6-11.9c-3,3.9-6,7.7-8.7,11.7C189.2,113.2,189.3,113.3,189.3,113.3z"
	pathData54 := "M204.9,128.9c0.3,1.9-0.6,2.9-2.7,3.1c0.7-1.6,1.5-2.6,2.9-3L204.9,128.9z"
	pathData55 := "M205.1,129.1c-0.2-1.3,0-2.3,1.6-2c0.3,0.1,0.9,0.5,0.9,0.5c-0.5,1.2-1.4,1.6-2.7,1.3    C204.9,128.9,205.1,129.1,205.1,129.1z"
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+hexcode)
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
	canvas.Path(pathData25, "fill:"+hexcode)
	canvas.Path(pathData26, "fill:"+hexcode)
	canvas.Path(pathData27, "fill:"+hexcode)
	canvas.Path(pathData28, "fill:"+hexcode)
	canvas.Path(pathData29, "fill:"+hexcode)
	canvas.Path(pathData30, "fill:"+hexcode)
	canvas.Path(pathData31, "fill:"+hexcode)
	canvas.Path(pathData32, "fill:"+hexcode)
	canvas.Path(pathData33, "fill:"+hexcode)
	canvas.Path(pathData34, "fill:"+hexcode)
	canvas.Path(pathData35, "fill:"+hexcode)
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

}