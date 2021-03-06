package heraldry

import svg "github.com/ajstarks/svgo"

func renderBearRampantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M53.2,170.3c-3.9,1.3-5.6,3.6-6.8,6.5c-0.7,1.8-1.8,3.7-3.9,2.9c-2-0.7-2.5-2.7-2.3-4.9c0.1-1.9,1.4-3.3,2.1-4.9    c0.5-1.2,0.8-2.1-0.2-3c-1.2-1.1-2.1-0.4-3.3,0.3c-2.4,1.4-3,4.1-4.9,5.8c-0.8,0.7-1.3,2.6-3,1.3c-1.2-1-2.1-1.8-1.3-3.9    c1.3-3.4,2.3-6.7,5.1-9.2c0.5-0.4,0.6-0.9-0.3-2.2c-1.5,1.8-2.9,3.5-4.3,5.1c-0.6,0.7-1.3,1.5-2.5,1.2c-1.5-0.4-0.8-1.6-0.9-2.4    c-0.4-3.1,0.7-5.6,3-7.8c5.4-5.3,11.9-7.7,19.2-8.4c4.7-0.5,9.3-1.3,13.4-3.5c7.1-3.8,14.9-4.4,22.5-5.8c8.6-1.6,17.3-3,25.9-4.7    c1.2-0.3,2.5-0.1,3.7-0.5c-1.3-3.9-4.6-7-8.2-7.5c-1.2-0.1-1.6-1.1-2.5-1.4c-1-0.3-1.8-1.9-3.3-0.7c-1.3,1.1-2,2.4-1.5,4    c0.2,0.8,1.8,1.6-0.1,2.2c-1.3,0.4-1.4-0.7-2-1.5c-2.7-3.2-3-6.4-0.4-9.7c1.8-2.3,1-3.8-0.9-5.4c-6.1-4.8-12.6-9.3-16.5-16.4    c-0.8-1.5-3.1-2-4-3.5c-3.1-4.6-7-8.4-10.9-12.3c-3.2-3.2-7.7-3.1-11.8-3.9c-2.3-0.4-4.8-0.1-7.2-0.1c-1.4,0-3.1,0.4-3.6-1.3    c-0.5-1.7,1.4-2.3,2.3-3.2c1.7-1.9,4.8-0.9,6.8-3.1c-5,0.3-9.3-2.4-14.1-1.9c-0.8,0.1-2.4-0.8-2.6-1.5c-0.4-1.3,1.2-1.6,2.2-2    c3.9-1.7,8.3-1.8,11.9-4.4c-2.4-2.9-4.9-5-9.2-3.1c-1.2,0.6-3.6,1-4.6-0.8c-1.2-2,1.2-2.5,2-3.6c3.5-4.8,8.7-2.6,13.3-3.2    c-0.3-2.9-2.5-3.9-4.7-5.1c-1.5-0.8-3.8-1.9-3-3.9c0.8-2.1,3.3-1.4,4.9-1c4.8,1.2,9,4.4,14.3,4.6c0.9,0,2.2,0.6,2.5,1.9    c1,4.5,4.5,6.8,7.9,9.3c5.6,4.2,10.7,9.4,16.7,12.9c6.5,3.8,12.9,7.7,18.9,12.1c5.9,4.3,11.9,8.5,17.7,12.8    c4.1,3,9.6,2.9,13.6,6.3c0.9,0.7,2.2-0.3,2.5-1.5c2.7-9.1,10.9-12.3,18.1-16.1c5.5-2.9,11.6-4.7,17.4-6.9c2.1-0.8,3.5-2,3.5-4.4    c0-2.9-2.2-2.4-3.9-2.4c-1.2,0-1.6-1-2.3-1.6c-0.9-0.9-1.8-2.2-3-2.1c-4.5,0.5-8.4-3.6-13.1-1.5c-0.9,0.4-2.3,0.2-3.4,0    c-2.4-0.4-3.1-1.7-1.4-3.7c1.5-1.7,3.1-3.1,5.3-3.9c1.7-0.6,2-1.8,0.5-3.1c-1.6-1.5-3.1-3.1-4.1-5.2c-2-4.3-0.1-7.5,4.8-7.7    c4.2-0.1,8.5-0.1,12.7,0c14.5,0.4,28.8-2.8,43.3-1.7c1.6,0.1,2.7-0.5,3.8-1.9c4.5-5.7,15.5-6.7,20.9-2c2,1.7,1.3,3.4-0.4,5    c-2,1.8-1.8,3.3,0.9,3.8c5.2,0.9,9.2,5.4,15,4.7c1-0.1,2.7,1.5,3.2,2.6c0.7,1.7-1.6,1.8-2.4,2.8c-2.3,3-6,4.5-7.9,8.3    c-1,1.9-3.8,5-7.4,4.8c-2.2-0.1-3.1,0.2-2,2.8c3.1,7.9-1.1,14.6-4.1,21.4c-1.9,4.2-4.5,8.1-6.9,12.1c-1.1,1.8-3.4,3.2-1.4,5.8    c0.2,0.2,0,1-0.3,1.1c-5.9,4-2.8,10-1.6,14.1c2.5,9,4,17.9,4.7,27.1c0.6,7.9,3,15.6,4.5,23.4c2.3,12.5,9.5,22.5,15.9,32.9    c4.3,6.9,6.6,14.6,8.4,22.5c2.1,9.2,4.8,18.3,5.7,27.7c0.6,6.4-0.7,12.7-2.6,18.8c-1.3,4.1-1,7.2,3.1,9.4c1.9,1,3.3,7.5,1.8,8.8    c-2.1,1.9-4.6,3.2-7.4,1.2c-6.4-4.2-13.4-7.9-15.7-16.1c-0.8-2.7,0-5.5-1.3-8.3c-1.1,3.7-2.5,7.3-3.2,11.1c-0.9,4.4-1,4.6-4.7,2.2    c-0.8,2.2-0.7,4,1.2,5.5c1.1,0.9,2.4,1.9,1.7,3.5c-0.6,1.4-2,1.4-3.3,1.3c-2.2-0.2-3.2,0.2-2,2.8c0.7,1.6,0.3,3.6,0.4,5.4    c0.3,13.9,0.6,13.1,9.3,19.8c-0.6,1.4-2.1,0.3-2.9,0.9c-0.9,0.6-1.9,1.3-2.3,3.1c2.9,0,5.7,0,8.4,0c0.2,0.3,0.4,0.6,0.6,1    c-4.4,1.6-1.9,6.7-5.4,8.8c-0.8,0.5-1.3,1.8-2.4,1.4c-5.3-1.7-9.3,1.8-13.7,3.5c-7.8,3-15.7,5.4-23.7,7.6c-3.5,0.9-5.5,4-8.3,6    c-1.5,1-2.6,2.5-3.9,3.9c-0.8,0.9-1.8,1.6-2.9,1c-1.1-0.5-1.2-1.7-1.3-2.9c-0.1-1.2,1-3-0.8-3.7c-1.5-0.6-2.1,1.1-3.1,1.7    c-2.1,1.2-3.8,2.8-5.3,4.6c-1.1,1.2-2.7,1.8-3.7,1c-1.2-0.9-0.5-2.4,0.2-3.7c1.5-2.9,2.7-6,4-8.9c0.4-0.8,0.6-1.8-0.1-2.3    c-0.6-0.4-1.6-0.2-2.3,0.4c-1.2,1.1-2.6,2-3.9,3.1c-1.1,0.9-2.1,2.2-3.5,0.9c-1.5-1.4-1.2-3,0.2-4.4c0.7-0.7,1.5-1.2,1.6-2.4    c0-0.4,0.8-0.6,0-1.2c-2.8-2.1-2.4-4.9,1.3-7c5.5-3.1,11.3-5.1,17.5-1.8c1.4,0.7,2-0.8,2.7-1.5c3.3-3.3,6.5-6.9,11.2-8.4    c1.7-0.6,0.6-1,0.2-1.5c-2.1-2.3-2.6-5-2.4-8c0.2-3.3-2.5-5.2-4.4-7.1c-4.5-4.5-7.8-9.4-8.4-16c-0.2-2.6-3-4-3.6-6.8    c-0.5-2.4-1-4.8-0.2-6.9c1.1-2.9,0-5.5-1.3-7.3c-1.3-1.8-3.6,0-5.5,0.7c-2.7,0.9-6.2,1.3-8.6-0.2c-3.4-2.1-4.2,0-5.6,2    c-0.8,1.2-0.7,3.6-2.4,3.4c-1.8-0.2-3-2.4-3.1-4.1c-0.2-2.9-1.1-3.2-3.3-1.8c-0.4,0.3-1.2,0.3-1.7,0.1c-5.4-2.5-9.7,1.2-14.5,2.5    c-1.3,0.4-2.1,1.3-2.8,2.4c-4.6,7.4-11.4,7.9-17.6,1.9c-4.3-4.2-9.1-7.9-13.6-11.8c-0.5-0.4-1.3-0.6-1.7-1.1    c-6.4-7.7-15.8-8.1-24.7-10.7c1.8-4.6,6.1-5,10.7-5.4c-2.3-3.5-5.1-5.5-9.4-5c-1.9,0.2-4.4,0.4-4.8-1.9c-0.4-2.5,2.4-2.4,4.1-3.2    c3.7-1.7,7.7-0.2,11.5-1c-2.3-5.8-8.5-9.4-13.9-7.9c-1,0.3-2,0.9-2.6-0.1c-0.6-1.1,0.3-2.1,1.1-2.9c3.9-3.6,10.2-4.4,14.8-1.7    c0.4,0.3,0.9,0.4,1.6,0.6c0-0.5,0.1-0.9-0.1-1.2c-0.6-1.4-3.8-1.2-2.9-2.9c0.9-1.8,4-1.2,5.3-0.6c4.9,2.5,10.2,4.8,13.3,9.8    c1.2,1.9,2,3.9,3.1,5.8c1.3,2,2.9,3.7,5.5,2c3.1-2,5.9,0.1,8.8,0.5c1.2,0.2,2.1,1.3,2.8-0.8c2.5-7.9,8.8-12.1,16-14.7    c8.8-3.3,17.6-6.7,27-7.8c1.8-0.2,3.6-0.8,2.1-3.9c-1.3-2.8-2.4-5.4-4.6-7.7c-5.3-5.5-10.5-11.2-16.3-17.4c-1.6,2-2.6,3.3-3.7,4.6    c-0.7,0.9-1.7,1-2.6,0.4c-1.1-0.7-1.5-1.9-0.7-2.9c5-6.3,0.7-10.8-2.6-15.7c-1.1-1.6-2.5-2.9-3.4-4.7c-1.5-2.6-4-0.1-5.7,0.7    c-2.2,1.1-4,2.2-6.6,1.2c-1.4-0.5-2.2,0.6-2.4,3.1c-0.1,1.5,0.5,3.4-1.4,4c-2,0.7-2.9-1.4-4.1-2.4c-0.6-0.5-1.1-1.3-1.4-2    c-0.8-2.5-1.7-3.6-4.8-2.4c-2.7,1-5.4-1.4-8.3-1.8c-3.7-0.6-5.3,0.5-4.4,4c0.3,1.2,0.8,2-0.4,2.8c-1.4,0.9-2.8,0.1-3.2-1.1    c-2.2-6-7.3-4.4-11.7-4.6c-2.8-0.1-5.5-0.7-8.3-0.9c-6-0.6-8.4,0.2-10.4,6c-1.1,3-3.3,5.4-4.3,8.3c-0.4,1.2-1,2-2.3,1.9    c-1.3-0.1-1.5-1.2-2.1-2.2C50.3,176.8,52.2,173.9,53.2,170.3z"
	pathData1 := "M144.3,254.2c-4.7,1.3-7.5,5.1-11.5,9.1c0.2-6.4,4.2-8.7,7.6-11.1c1.7-1.2,3.8-1.9,5.9-2.8    c5.7-2.4,10.1-6.8,15.1-10.4c2.8,0.2,5.2-1.5,8-1.7c-2.7,0.3-5.2,1.6-7.9,1.8c-3.3,0.1-6.4,1.8-8.4,3.7c-3.9,3.7-8.8,5.5-13.1,8.3    c-3.1,2-6.2,3.9-8.4,6.9c-0.8,1.1-1.8,2.4-1.3,3.8c1.7,5-1.8,8.5-3.5,12.4c-1.2,2.7-7.4,2.8-9,0.4c-0.7-1.2,0.5-1.2,1-1.6    c0.7-0.6,2.2-1.1,1.5-2.3c-0.8-1.6-1.8-0.4-2.8,0c-4,1.8-6.6-0.6-9.8-3.1c1.5-0.2,2.5-0.3,3.4-0.4c-6-1.8-10.1-5.9-14-10.6    c2.3-1,3.5,1.2,5.9,1.5c-0.6-3.4-3.6-3.2-5.3-3.3c-2.5-0.1-4.9-0.9-7.4-1.2c-1.5-0.1-4-0.7-3.5-3.2c0.4-2.1,1.5-3.7,4.3-3.5    c2.2,0.2,4.4,0.6,6.4,1.4c2,0.7,3.4,0.2,5.1-1c-5.1-1-10-2.1-14.9-2.7c-2.6-0.3-3-2.1-3.5-3.9c-0.4-1.4,0.3-3.1,1.4-3.6    c1.7-0.7,3.3,0,4.9,1.3c3.4,2.7,7.6,3.9,13.1,4.7c-5.2-3.3-10.1-4.7-14.2-7.7c-2.2-1.6-2.4-3.9-3.3-5.8c-0.6-1.4,1.3-2.3,2.4-3.1    c1.1-0.9,1.6,0.6,2.1,1.2c5,6.4,11.6,10.1,19.4,12c1.8,0.4,3.5,0.5,6.2,0.5c-2.3-0.8-3.8-1.3-5.5-2c1-0.4,1.7-0.8,2.5-1    c2.2-0.4,4.6,0,6.1,1.5c1.3,1.4-0.6,2.8-1.4,4.1c-1.5,2.4,0.1,10.9,2.3,12.4c0.7,0.5,1.9,0.9,2.2-0.7c0.9-5.5,0.9-5.6,4.2-8.5    c-2.4-0.4-2.4-0.4-5.3,2.5c-0.5,0.5-1,1.3-1.7,0.9c-0.9-0.5-0.6-1.4-0.3-2.2c2.2-5.7,4.4-11.3,11.1-13.2c1.2-0.3,2.1-1.5,3.2-2.4    c-2.2-2.2-3.2,0.8-4.8,1c1.6-4.7,2.4-5.4,13.4-10.1c5.8-2.5,12-0.2,18.3-1.4c-1.7-0.9-3.1-1.7-4.6-2.5c3.6-3.2,6.9,1.5,10.6,0.4    c1.3-0.4,1,2,1.5,3c0.7,1.5,1.3,0.4,1.8,0.3c1.5-0.3,0.2-0.9,0-1.3c-0.2-0.5-0.5-1-0.8-1.4c-1.8-2.1,0-3.2,1.4-3.8    c1.9-0.8,2.6,0.1,1.5,2c2.3,1.3,3.8,5,6.6,4c2.3-0.8,3.8-3.6,3.1-6.7c-0.2,0.1-0.5,0.2-0.6,0.4c-0.9,1.5-1.3,4.2-2.7,4.1    c-2-0.2-2.5-2.9-3.4-4.7c-1.8-3.8-0.1-8-1.6-11.9c-0.4,1.1-2.2,1.9-0.4,3.6c0.6,0.6,0.7,1.8-0.2,2.2c-0.9,0.4-1.6,0.2-2.2-1.1    c-2-4.7-6.1-8.5-2.9-14.8c1.1-2.2-0.1-6.7-0.7-11.3c-2.6,5.5,0.3,11.5-5.4,15.3c1.4-7.1,1.2-14,2.6-21.9c-2.3,2.4-2.9,4.3-2.9,6.4    c0,2.9-1.1,3.7-3.9,2.1c1.2,1.4-0.7,3.8,2.1,4.6c0.9,0.2,1.1,1.3,0.9,2.2c-0.3,1.2-1.3,0.8-2.1,0.8c-0.9,0-2.1,0.5-2.1-1.2    c-0.1-2.6-1.8-4.7-1.3-7.8c0.3-2-1.1-5-0.3-7.6c0.3-1,0.7-2.5,1.4-2.7c1.3-0.4,1.5,1.1,2.2,2c0.5,0.7,0.2,2.3,1.7,2    c1.2-0.3,1.7-1.4,1.4-2.7c-0.4-2.1,0.7-4,0.9-6.1c-1.8-0.1-1.4,1.2-1.8,2.1c-0.4,0.9-0.3,2.5-1.8,2.4c-1-0.1-1.3-1.2-1.7-2.1    c-0.9-1.9-1.4-3.8,0.1-5.3c1.9-1.8,4.3-0.3,6,0.5c2.1,1.1,3.4,3.4,6,4.3c1.8,0.7,1.7,3.5,1.7,5.7c0,1.2-1,3.5,0.8,3.6    c2.4,0.1,0.8-2.1,1-3.3c0.6-3.2,1.8-6.3,1.4-9.6c-1.4-0.1-1.4,1.1-2,2c-0.8-1.3-1.5-2.5-2.2-3.7c-1.6-0.2-1.4,1.4-2.2,2.1    c-2.9-2.6-2-6.9-4.6-10.3c0,3,0,5.4,0,8.1c-3-1.3-4.5-3-4.2-6.1c0.2-2.2,0.1-4.5,0-6.8c0-0.8,0.7-2.3-0.8-2.4    c-1.3-0.1-1.9,1.1-1.7,2.4c0.3,1.7-0.9,2.9-1,4.4c0,0.7-0.6,1.8-1.6,1c-0.5-0.4-2.2-0.9-0.9-2.1c1.5-1.4,1-3.1,0.8-4.9    c-4.4,1.2-5.8,4.8-7.4,8.3c-0.3,0.7,0.6,1.2,1.1,1.6c0.4-0.5,1.2-0.9,1.3-1.4c0.1-1.1,0.4-2,1.3-2.4c0.2-0.1,0.9,0.6,1.1,1    c0.5,1.2-0.4,2.3-1.1,2.9c-1.8,1.7-3.4,3.4-4.7,5.5c-2.3,3.6-2.4,3.6-1.8,6.8c-2.3-2.6-2.3-2.6-1.6-5.8c0.3-1.2-0.3-2.3,2-1.8    c1,0.2,1.8-2.6,1.1-4.5c-2.2,0.3-3,2.3-4.4,3.6c-1.4,1.2-2.5,2.7-3.8,4.2c-2.7-4,0-6.8,1.7-10.1c-5.2,3.1-8.9,8.9-15.9,9.1    c1.3-2.6,4.8-2.1,5.7-5.2c-3,0.5-5.4,2.5-8.1,3.4c-3.8,1.3-1.7,5.5-3.8,7.5c-3.3-3-3-5.7,0.5-8.3c1.9-1.4,2.1-3.2,0.1-5.2    c-1.4,4.9-4.9,6.9-9.7,7.7c0.1-0.7,0-1.1,0.1-1.3c0.8-1.3,3-2.6,1.6-3.9c-2.1-2-2.8,1.1-4.1,2c-1.6,1.2-2.9,2.5-5.4,2.1    c1.4-1.9,3-3.4,4.9-4.6c-7.2,0.3-11.5,4.1-12.1,10.9c-0.7-6.8-5-6.4-9.5-5.1c-0.3-1.1-0.1-1.6,0.3-2c0.5-0.6,1.6-1.2,0.8-2    c-0.8-0.8-1.7-0.3-2.6,0.4c-1.4,1-2.6,2.5-4.5,2.7c-0.9,0.1-2.9,0.4-0.5-1.2c0.2-0.1-0.7-1.3-1.6-1.2c-0.8,0.1-1.6,0.4-2.3,0.8    c-1.3,0.6-2.8,1.5-3.9,0.5c-1.5-1.4,0.8-2.4,1.2-4.2c-4.8,2.2-7.9,5.3-10.3,9.3c-0.4,0.6-0.3,2.9-2.1,1.6    c-1.1-0.8-2.3-1.5-1.2-3.7c1.3-2.8,3.4-5,6.2-8.1c-4.8,1.2-6.4,4-7.7,7.1c-1.1,2.4-2.6,3.7-5.6,2.7c-1.3-0.5-3.4-0.7-3.5-1.6    c-0.1-1.4,1.4-2.6,2.5-3.8c2-2,5.5-2.1,7-6.2c-4.1,2.6-7.4,4.7-10.7,6.9c-1,0.7-1.7,1.8-2.7,0.3c-1-1.4-3.4-2.1-3-4.3    c0.3-1.4,1.9-1.7,3.3-1.8c3.4-0.1,5.2-2.7,7.6-4.4c-3.2,0.9-5.9,3.4-9.5,2.6c-0.8-0.2-1.9,2-2.5-0.3c-0.4-1.5-0.2-2.4,0.9-3.5    c3.3-3.1,7.6-3,11.6-3.7c5.3-0.8,10.5-1.8,15.3-4.3c6.8-3.6,14.5-3.8,21.9-5.1c5.4-0.9,10.9-1.5,16.2-2.8    c5.1-1.2,10.5-1.3,15.3-3.8c1,1.8-2,2.1-0.9,4.2c3.6-5.2,9-6.2,14.6-7c1.1-0.2,2.2-1,3.3-1.6c0.7-0.4,1.3-1.5,2-0.3    c0.4,0.6-0.1,1.3-0.6,2.1c-2.8,4.6-6.7,7.9-11.5,10.3c-1.5,0.7,0.7,2.1-0.2,3.1c0.6,0.6,1.2,1.5,2.1,0.7c1-1,2.2-0.1,3.2-0.9    c-0.6-1.3-2.7,0.4-2.9-1.3c-0.2-1.4,1.3-1.5,2.2-2.1c4.6-2.7,7-7.4,10.2-11.3c1-1.3,2.1-3.4-0.5-4.5c-3.1-1.3-1.3-3.4-0.1-3.8    c3.8-1.3,5.7-5.7,10.3-5.9c1.1-0.1,1.9-3.4,3.9-4.5c-4.5,0.1-8.5,2-10.6,5.5c-1.4,2.3-5.1,0.8-5,4.1c-0.7-0.4-1.6-0.6-2.2-1.2    c-1.3-1.1-1-7.1,0.2-8.3c0.8-0.8,1.8-0.6,2.2-0.1c0.7,0.8,2.6,1.2,1.5,2.9c-0.5,0.8-0.6,1.5-0.1,2.7c1.7-2.6,2.5-5.7,5.2-7.6    c-1.2,0-2.4-0.1-3.6,0c-1.4,0.1-3.2,2-3.9-0.1c-0.3-0.8,1.7-3.5,2.7-3.5c4.1-0.1,7.2-3.3,11.5-3.2c0,1.4-2.2,2.1-1.1,3.6l-0.2,0    l0-0.1c2.1,0.6,3-1.9,4.9-1.8c0.2,0,0.7,0.1,0.6-0.7c-0.2-3-0.1-3,2.4-4.6c-5.7,0.9-11.3,1.9-17,2.8c8.1-6.8,18.1-9.2,27-13.5    c-0.9-0.1-2.1-0.2-3.5-0.4c2.3-1.7,4.5-2.8,8.2-2.2c-4.2-2.2-7.1-2.8-9.7,0.7c-1.2,1.7-3.2,2.8-5.4,2.4c-2.7-0.5-2.2,2.3-3.6,3.1    c-3.4,1.9-6.9,3.6-10.8,5.6c1.8-5.3,6.7-6.7,10.2-9.4c3.3-2.6,7.1-4.6,10.7-6.8c1.3-0.7,2.6-1.4,4.3-0.6c2.2,1.1,2.5,0.8,2.4-1.9    c1.7-0.3,1,1.8,2.1,2.1c1,0.3,0.4,2.2,1.8,2c1.4-0.3,1.4-1.3,0.8-2.4c-0.1-0.1,0.3-0.7,0.6-0.8c3.6-1.5,4.9-4.9,6.7-7.9    c0.8-1.2,2.1-1.5,3.5-1.8c-6.9-7-7.4-8.3-3.5-10.3c-1.5-0.9-3.5-0.2-4.3-2.5c-0.5-1.5-2-4.1-4.5-0.9c-1,1.3-2.6-0.8-4-1.1    c-2.4-0.5-4.8-0.9-7.1-1.9c-4.6-1.9-5-2.7-3.8-9.8c-0.5-0.6-0.7,0.1-0.9,0.2c-0.7,0.8-1.6,1.2-2.6,0.9c-0.2-0.1-0.5-0.7-0.5-1    c0.4-1.2,1.4-1.5,2.7-1.5c8.5,0,17,0.4,25.5-0.1c10.5-0.5,20.9-2.1,31.4-1.6c4.9,0.2,9.6,2.3,11.6,5.5c-1.3,1.5-3.3-0.6-5,0.5    c3,2.4,6,1.9,8.4,0.3c6.2-4,11.7-1.9,17.3,1.3c1.1,0.6,2.1,1.2,3.2,1.8c0.6,0.3,1.9-0.5,2,0.6c0,0.6-0.1,2.1-1.6,2    c-2.1-0.2-1.8,2.3-3.1,2.9c-3.5,1.8-3.7,8.5-9.6,6.2c-0.5,1.3,0.8,1.6,0.9,2.6c-2.4,0.5-4.8-0.1-7,2.2c-1.5,1.5-4.4-1-6.4-2.4    c-1.9-1.3-1.9-4.1-4.4-4.7c-0.8-0.2-0.5-2.9-2.1-1.1c-0.6,0.7-0.9,1.8-0.1,2.9c3.1,4.1,6.6,7.6,12.3,7.6c1.5,0,3.1-0.7,4.5,1.1    c-1.5,1.3-3,2.1-4.9,1c-1.3-0.8-1.8,0.1-2.3,0.9c-0.5,0.8,0.1,1.1,0.9,1.3c8.1,2.2,8.1,2.2,3.3,9.9c-0.9-2.8,1.4-6.7-4.1-8.6    c2,4,3.3,7.6,0.7,11c3.1,2.4,1,4.7-0.2,7c-1.3,2.4-2.3,5-4,7c-1.3,1.5-1.7,4.2-3.5,4.6c-2.3,0.6-1.3-3.7-4-3.7    c0.6,2.9,0.9,5.7,1.7,8.3c0.3,1,0.4,1.5-0.3,2.1c-1.1,0.8-1.5-0.3-1.7-0.9c-0.5-1.8-2.3-1.3-2.6-0.9c-0.8,0.8,0.2,1.6,0.8,2.4    c1.9,3,1.4,6.8,2.6,10c2,5.8,0.2,11.6-0.2,17.7c3.4-2.6,1.8-6.9,4-9.7c1.8,6.5,2.8,12.9,3,19.4c0.2,6.5,1.9,12.7,3.5,19.3    c-2.2-1.1-3.1-2.8-4.2-4.3c-0.6-0.8-1.2-3.1-2.9-1.7c-1,0.9,0.3,2,0.9,3c2.1,3.4,5.5,5.8,6.9,9.7c0.5,1.4,1.4,2.7,0.2,4.5    c-1.3-1.4-2.5-2.8-3.8-4.1c-0.5-0.5-1.1-1.6-2-0.9c-0.7,0.6-0.6,1.7,0.1,2.3c2.9,2.2,4,4.8,3,8.5c-0.5,2,2.1,3,2.8,4.8    c1.1-3.2-1.1-6.6,1-9.8c2.4,2.8,2.8,6.6,4.6,9.7c5.8,10,12.3,19.7,16.3,30.6c3,7.9,3.9,16.4,6.5,24.5c0.2,0.5,0.2,1.4-0.1,1.6    c-3.4,2.6-2.5,7.1-5.2,10.9c4.7-1.1,4.4-5,6.4-7.7c-0.8,4.7,3.8,9.1-0.1,13.9c-0.7,0.8-1,1.9-1.4,2.9c-0.4,0.9-1,2.1,0.4,2.4    c1.6,0.4,0.4-1.7,1.6-2.1c0.8,3.4-0.9,6.4-1.7,9.5c-1.5,5.6-0.9,8.2,3.9,11.1c2.3,1.4,1.7,3.1,0.8,4.5c-1.3,2-3,0.2-3.7-0.6    c-2.1-2.3-3.5-5.2-4.2-8.3c-0.2-1-0.5-2-1.6-1.6c-1.1,0.4-0.7,1.3-0.5,2.4c0.4,2.3,2.9,3.7,2.8,6.9c-3.7-4.2-6-8.3-4.3-13.4    c0.6-1.8-1.1-2.8-0.8-4.5c-2.3,4.4-2.9,8.6,0.3,12.9c-5.3-3.9-7.6-13.2-4.8-18.4c3.7-6.9,7-13.9,7.7-21.7    c0.4-3.7-0.2-7.6-2.2-11.1c-0.6-1-1.3-1.2-1.9-0.9c-1.1,0.4,0,1.2,0.2,1.6c3.5,6.1,0.1,12.5,0.7,18.7c0,0.4-0.5,1.2-0.9,1.3    c-1.5,0.5-1.2-1.4-2.2-1.9c-2.1,4.2-2,9.2-5,13.7c-0.6-2.3-0.3-4.3-2.4-5.7c-1.5,7.5-2.9,14.7-4.4,22.1c-2.5-1.7-0.6-4.3-1.9-7    c-0.2,3.7-4.9,3.4-4.5,7.1c0.3,3-1.4,6.8,4,8.3c-4.2,0.9-7,0.1-9.7-3.3c-0.4,4,2.9,4.5,3.3,7c0.5,3.2,0.7,6.3,0.7,9.5    c0,6.3-0.1,12.6,6.5,16.2c0.3,0.2,1.9,1,0.1,1.4c-1,0.2-1.3,2.6-2.9,0.8c-3.6-4.3-8.6-7.1-11.4-12.1c-0.6-1-1.3-2.2-2.9-2.2    c-1.6,2.8,0.6,5,2.1,6c3.3,2.3,5.8,5.2,8.1,8.4c2.9,4,6.6,6.4,12.3,4.8c-1.3,1.9-2.8,0.7-4.1,1.6c1.9,0.4,3.5,1.5,3,3.3    c-0.4,1.6-2.1,2.5-4.1,2.3c-3.5-0.5-0.6-3.3-2.4-4.9c-4.2,4.7-10.6,6.7-15.4,12c-0.9-2.7,1.1-3.7,0.5-5.5    c-2.4,2.8-3.6,6.5-7.8,6.9c-3.5,0.3-6.8,1.6-10.1,3.6c0.3-4.4,5.6-4.6,6-8.8c-1.7,0.9-3,1.8-4.3,3.3c-2.9,3.3-6.2,6.3-9.3,9.5    c-0.8,0.8-1.5,0.8-1.9-0.1c-0.4-0.9-2.8-1-1.2-2.8c1.9-2.2,3.8-4.5,6.4-7.6c-4.6,1.7-7,4.4-9,7.5c-2,3.1-4.4,3.3-6,0.8    c-1.7-2.6,0.2-3.5,1.9-4.5c1.4-0.9,3-1.3,4.3-2.3c1.2-0.9,2.1-2,3.1-3.3c-2.7,1.4-5.5,2.7-8.1,4.3c-1.9,1.2-3.9,1.2-5-0.2    c-1-1.3-2.2-3.3-1.4-5.2c0.8-2,2.8-1.8,4.7-1.9c2.9-0.2,5.2,2.6,8.2,1.5c-0.5-2.1-2.5-1-4.7-2.5c5.9,1.3,7.6-3,10.4-5.2    c4.7-3.7,10-5,16.7-6.2c-5.7-1.1-10-2.6-10.3-9.2c2.9,3.5,6.1,6.4,9.9,5.9c4.1-0.6,6.6,3.7,11.1,2c-7.1-3.1-13.5-6.5-16.9-13.7    c1.3-1.1,2.9-0.2,4.4-1c-3.5-1.2-6-3.7-9.5-6.2c1.9,6.9,4.4,12.3,9.3,16.5c-3.5-1.4-6.5-3.9-7.9-6.9c-1.7-3.6-4.2-6.2-6.8-8.9    c-3.3-3.6-6.3-7.1-6.6-12.5c-0.1-2.2-2.4-4.2-3.1-6.5c-1.9-5.9-0.2-11.2,3.4-16.1c-3.1-1.9-3.1,1.8-5,2.8    c-0.9-7.4,2.5-12.9,6.7-18.7c-4,0.6-4.3,4.3-6.6,6.1c-1.8-4.6,0.4-8.6,2.7-11.6c6.5-8.3,12.9-16.6,16.8-27.4c0.9,3.2-2,5.2-0.5,8    c3-5.2,4.4-10.3,3.4-15.7c-4,5.9-7.9,12-10.9,18.2c-2,0.5-4.7-3.2-5.1,1.7c0,0.4-1.4,0.7-2.1,1.1c-0.1-0.8-0.5-1.8-0.2-2.5    c0.7-1.8,2.7-2.6,4.2-3.1c6.5-2.4,7.6-8.4,9.2-13.8c0.5-1.9,1-4,1.2-6.1c0.1-1,1.3-2.5,0.2-2.9c-1.4-0.6-2.1,0.9-2.1,2.3    c0,0.3-0.4,0.6-0.9,1.3c-2.1-4.4,0.6-8.3,1-13c-4,4.6-4.2,11.5-2,15.5c1.8,3.1-0.3,6.4-1.5,9.4c-0.4,1.1-1.4,2-2.1,3    c-0.5,0.7-0.4,2.6-2.2,1.2c-1-0.8-0.8-1.5,0-2.4c1.9-2.1,3-4.1,1.1-7.1c-0.9-1.4,0.6-3.4,0.8-5.1c0.3-1.9-0.5-3.3-0.4-5.4    c-2.6,2.6-1.2,5.4-1.5,7.9c-0.3,2.3-0.1,2.4-2.7,2.2c-0.9-0.1-0.7,0.5-0.8,0.9c-1.2,4.7-5,8-6.9,12.9c1.6-0.7,2.8-1.1,4.5-1.8    c-0.8,3-0.9,5.7-4,8.3c3.2-5.2-2.5-5.2-3.6-7.9c0.8,2.6,1.6,5.2,2.4,7.8c-1.2,1.1-2.7,0-4,0.9c0.9,0.4,1.8,1.1,2.6,1    c1.1-0.2,2-0.3,2.2,0.9c0.3,1.8-1.3,0.9-2.1,1.1c-1.7,0.4-3.4,0.8-5.1,1.5c2.4,0.1,4.8,0.2,8,0.4c-3.4,3.3-7.5,1.8-10.7,3.7    c3.4,1.2,6.7-1.8,10.3,0.2c-2.5,1.8-5.9,0.7-7.9,3.2c2.3,1.3,4.2-1.1,6.4-0.5c0.6,0.2,1.5-0.3,1.8,0.6c0.3,1-0.5,1.8-1.3,1.9    c-2.7,0.5-5.1,1.8-7.7,2.6c-9.4,2.9-18.4,6.3-25.1,14.1c-1.2,1.4-3.9,1.6-6.2,1.9c9.9-7.1,19.2-15.1,30.8-19.6    c-3.4-1.3-5.9,1.3-8.5,2.3c-2.6,1.1-5.5,2.6-7.8,4.6c-2.5,2.1-4.9,4.2-8,5.2c-3.1,1-5,3.8-7.9,4.9c-2.6,1-2,6.3-7.1,5    c3.2-3.9,5.4-7.9,9.6-10.2C143.5,256.3,145,255.7,144.3,254.2c2.2-0.2,4-1.6,6.2-2.1c2.3-1.4,4.7-2.7,6.1-5.2    c1.8-0.7,3.5-1.7,5.2-2.6c1.4-0.4,2.6-1.3,3.1-1.5c-0.5,0.3-1.9,1-3.3,1.6c-1.7,0.8-3.3,1.7-5,2.5c-1.4,2.4-3.9,3.7-6.1,5.2    C148.2,252.5,146.5,253.9,144.3,254.2z"
	pathData2 := "M92.3,106.9c1.2-0.1,2.5-0.1,4.5-0.3c-3.1-3.3-3.5-7-3-10.8c-2.7,0-0.7,3.1-3.2,2.8c-0.9-2.1,0.5-4.8-0.8-7.3    c-1.6,3.3-2.2,6.1,0.5,13.3c-8.3-8.9-16.5-17.8-24.9-26.9c1.7-0.7,2.7-1.3,3.7-1.4c1.4-0.2,2.5-0.4,2.8,1.9c0.3,2.2,1,4.7,4.1,5.9    c-0.7-4.3-4-7.8-1.8-12c0.3-0.5-0.8-1.2-1-0.8c-3.5,4.7-9.2,1.2-13.3,3.7c-0.8,0.5-1.7,0.8-2.3,0.2c-1.4-1.3-5.1,0.1-4.4-3.6    c0.7-3.5,4.2-5.8,7.4-5c1.4,0.4,2.7,0.5,4.7-0.2c-4.1-1.8-7.8-1.4-11.4-0.7c-1.2,0.2-2.5,0.4-3.7,0.4c-2.2,0-3.5-1.8-3.5-3.5    c0-1.6,2-1.6,3.5-1.6c4.9,0.3,9.9-1.2,14.6,1.5c-2.9-3.9-7.1-3.5-11-4c-1.9-0.2-3.9-0.5-4.7-2.1c-0.8-1.7,1.2-3.1,1.8-4.7    c0.3-0.7,1.1-1.4,2-0.5c3.2,3,7.5,3.3,11.4,4.2c4.4,1,8.4,2.6,11.6,5.6c9.8,9.2,21.9,15,32.8,22.6c7.3,5.1,13.5,12,22.7,14.2    c1.8,0.4,3.8,2.3,4.6,4.3c0.9,2.2-1.2,3.4-2,5.1c-1,2.2-2.7,4.5-1.8,7.3c0.6,2-2,2.8-2,4.5c0.1,1.6-1,2.8-0.6,4.7    c0.2,0.9,0.3,3-0.9,4.6c-1.7-1.7-4.8-2.8-0.9-5.4c1-0.6,0.7-1.4,0.7-2.2c0-4.7,1.5-9.2,2.1-14.7c-1.4,2.3-3,3.4-3,5.8    c0,1.8,0.1,4-2,5.2c-0.6,0.4-0.7,1.3,0.2,1.6c1.5,0.4,1.3,1.6,0.7,2.2c-0.9,1-1.1-0.6-1.6-0.8c-1.2-0.4-1.7,0.8-1.3,1.6    c0.4,0.8,0.8,1.4,0.6,2.3c-4-0.7-3.1-3.7-3.4-6.4c-0.3-2.8-0.6-5.1,0.7-8c1.2-2.9,0.7-6.8-1.4-11.2c0.2,5.1-2.8,6.4-5.7,7.9    c-2.3,1.1-2.2,3.2-1.6,6.3c0.9-2.6,2.9-3.3,4.1-4.7c0.5-0.6,1.1-1.4,1.9-0.3c0.5,0.7,1.8,1.6,1.2,2c-4.6,3.2-0.8,7.5-1.7,11.3    c-1.1-0.3-0.9-1.9-2.6-2c1.8,1.9-0.1,5,2.3,6.3c1.6,0.9,1.9,2,0.8,2.8c-2.4,1.8-0.5-1.1-1.4-1.5c-3.1-1.3-3.1-4.1-3.1-6.9    c0-1.1,0.3-2.3-1.2-3.2c-1,2.7-1.9,5.6-5,6.5c-1.8,0.5-2.8-0.8-3.8-2.2c1.3-0.4,2.6-0.8,3.8-1.3c2.8-1.1,3.8-4.7,1.8-7    c-0.3-0.3-1-0.4-1.5-0.4c-0.7,0-0.8,0.6-0.8,1.2c-0.1,4.3-3,6-6.7,6.3c-2.3,0.2-1.9,2.3-3.5,3.6c0.4-4.2,3-6.9,5.1-9.4    c1.7-2.2,3.4-3.5,2-6.6c-0.6-1.4-0.1-3.3-0.1-6c-2.9,3.5-1.1,6.8-2,9.7c-2.4-0.9-2.1-3.5-2.9-5.3c-0.9-1.9,0.3-4.3-1.2-6.2    c-1,4.2-1.3,8.4,1.3,12.8c-3.4-0.7-4.1-4.1-7.2-4.1C92.1,107.6,92.2,107.3,92.3,106.9z"
	pathData3 := "M173.2,254.3c-0.1,3-3.8,1.8-4.6,4.3c2.4,0.8,4-2,6.3-1.6c1.1,0.2,2.7-2.1,3.4,0c0.7,2,2.5,4.7-0.1,6.5    c-2.3,1.6-5,2.6-7.7,3.7c-0.9,0.4-1.6-0.4-2-1.1c-0.1-0.2,0.5-1.1,0.9-1.1c3.6-0.6,5.2-4.4,9.5-5c-3.5-0.8-5.9,0.3-8.3,1.2    c-3.4,1.2-6.3,3.6-10.1,4.2c-2.4,0.4-3.9,2.4-4.5,4.9c-1.2-5.9,9.1-15.4,17.2-15.8L173.2,254.3z"
	pathData4 := "M142.9,186c0.6-4,4.6-6.8,1.2-10.5c3.4-2.6,1.2-6.9,3-9.9c0-0.1,0.4-2.1,1.8-1.3c0.9,0.5,0.9,1.4,0.8,2.2    c-0.1,0.8-0.6,1.5-0.8,2.3c-0.1,0.5,0.1,1.3,0.4,1.5c0.9,0.5,1.2-0.2,1.2-1c0-2,1.5-1.5,2.4-1.3c1.9,0.4,0.1,1.5,0.1,2.2    c0.2,5.7,0.1,11.4,0.1,17.4c-7.2-3.9-1.2-9.6-2.8-14.4C148.5,177.8,147.2,182.5,142.9,186z"
	pathData5 := "M237.2,34.4c-1.3,3.1-3.1,5.8-6,7.5c-2.4,1.4-3.1-1.1-4.5-2c-1.2-0.7-2.4-1.7-3.6-1.8    c-4.2-0.4-2.4-2.4-1.1-3.9c2.2-2.7,12.8-3.3,14.9-1C237.1,33.5,237.1,34,237.2,34.4z"
	pathData6 := "M178.6,248.4c-1.2,2.4-3.2,2.8-4.9,3.4c-7.2,2.9-14.5,5.7-19.8,11.7c-1.6,1.8-3.6,0.9-6.3,1.4    C156.5,255.7,166.9,251.3,178.6,248.4z"
	pathData7 := "M163.6,59c3.6-2.6,7.1-3.9,11.3-1.4c1.4,0.8,2.6,1.7,3.4,2.9c0.9,1.4,0.4,2.9-1,4c-1.5,1.4-1.9-0.4-2.5-0.9    c-1.8-2-4-2.8-6.7-2.5C166.2,61.2,164.6,60.8,163.6,59z"
	pathData8 := "M66.3,54.1c-2.8-0.8-5.6-1.5-8.3-2.4c-2.1-0.7-3.4-2.1-3.4-4.5c0-1.6,0.9-2.3,2.4-2.3c2.3,0,4.3,0.4,5.3,2.9    C63.4,50,63,53.2,66.3,54.1z"
	pathData9 := "M103.7,235.6c-3.7-1.8-5.9-3.5-7.7-5.6c-0.9-1.1-1.9-2.3-0.1-3.7c1.8-1.5,2.9-0.2,3.6,1    C101,229.6,102,232.1,103.7,235.6z"
	pathData10 := "M166.9,55.4c-4.4,1.8-7.8,5.2-12.7,4.6C157.8,55.4,162.6,53.5,166.9,55.4z"
	pathData11 := "M73.6,224.6c3.6-1.7,7.2-4.1,10.8-0.2c0.5,0.6,2.3,0.2,1.5,1.8c-0.6,1.1-1.6,1.4-2.9,1    C79.9,226.1,77.2,224,73.6,224.6z"
	pathData12 := "M154.9,157.6c2,0.6,2.6,2.2,2.3,3.5c-0.4,1.8-3.3,2.4-2.7,5c-0.4,0-1,0.1-1.2-0.1c-0.9-1-1.9-1.9-1.2-3.8    C152.8,160.4,153.1,158.5,154.9,157.6z"
	pathData13 := "M47.5,53.6c-3.9-1.9-7.5-1-11.1,0.4c2.1-2.8,4.7-4.6,8.4-4C46.3,50.2,46.3,50.2,47.5,53.6z"
	pathData14 := "M28.3,162.4c0.4-4.3,2.6-6.4,5.4-8c0.7-0.4,1.3-1.2,1.6,0.4c0.1,0.8,1.6,1.5-0.2,2.1    C32.4,157.8,30.9,160.1,28.3,162.4z"
	pathData15 := "M177.2,346c2.2,3.7,1,5.9-5.5,9.7C173.6,352.2,175.3,349.3,177.2,346z"
	pathData16 := "M165.7,345.5c1.5-1.9,2.9-3.8,4.4-5.6c0.7-0.9,1.6-1.7,2.8-1.2c0.8,0.3,0.5,1.1,0.6,1.8    c0.1,1.3-0.5,2.2-1.8,2.4C169.3,343,167.9,344.6,165.7,345.5z"
	pathData17 := "M31.7,170.8c0.9-2.1,1.8-4.1,2.5-6.3c0.5-1.6,1.9-1.8,3.1-2.2c0.3-0.1,1.3,0.7,1.5,1.3    c0.7,1.9-1.4,1.4-2.1,2.2c-1.5,1.8-3.1,3.5-4.6,5.2C32,171,31.9,170.9,31.7,170.8z"
	pathData18 := "M73.4,238.9c2.4-2.4,5.2-2.3,8.1-2.2c1.7,0.1,1.1,1.3,1.1,2.2c0,0.8-0.4,1.6-1.2,1    C79,238.4,76.2,239.9,73.4,238.9z"
	pathData19 := "M49.8,169.9c-3.6,1.5-4.3,4.9-6.3,7.1c-0.5-1.7,0.9-7,2-7.9C47,168,48,169.7,49.8,169.9z"
	pathData20 := "M43.5,74.6c1.8-2.4,4-2.1,5.8-2.9c1.5-0.6,2.4-0.4,2.5,0.8c0.1,0.8-0.4,2.4-2.2,2.1    C47.7,74.4,45.9,74.6,43.5,74.6z"
	pathData21 := "M85.8,247.8c-1.4,6.2-5.4,1.2-8,2.1C80.2,248.3,82.7,247.5,85.8,247.8z"
	pathData22 := "M87.7,219.2c3.3,1.9,5.6,3.5,7.4,5.8C92.6,226.4,90.4,224.8,87.7,219.2z"
	pathData23 := "M54.1,178.6c-1.2-2.7,0.8-5.2,0.6-7.9C58.4,172.4,58.3,173.3,54.1,178.6z"
	pathData24 := "M44.9,40.8c3.5-0.4,5.2,1.7,7.5,2.5c1.6,0.5,0,1.6-0.3,2.3c-0.6,1.9-1.2,0.1-1.5-0.2    C48.9,44,47.7,42.1,44.9,40.8z"
	pathData25 := "M131.7,125.3c-0.3-0.7-0.6-1-0.6-1.4c0-0.8,0.1-1.6,0.2-2.4c1.2,0.2,2.3-1.5,3.6-0.2c0.7,0.7,1.5,1.2,0.4,2.2    c-0.3,0.2-0.4,0.6-0.5,1C133.8,122.9,132.8,122.6,131.7,125.3z"
	pathData26 := "M166.4,337.5c1.2-2.6,2.9-3.4,4.7-4c0.4-0.2,1-1.3,1.5-0.1c0.1,0.4-0.1,1.4-0.3,1.5    C170.3,335.3,169,337.5,166.4,337.5z"
	pathData27 := "M46.1,66.3c-2.8-0.9-5.3-1.2-8.1-1C41.5,62.6,43.5,62.8,46.1,66.3z"
	pathData28 := "M186.6,348.9c4.9,3.2-0.8,4.2-0.8,6.3C185.1,353,185.2,350.9,186.6,348.9z"
	pathData29 := "M173.4,254.4c1.1-1.9,3.2-1.6,4.9-2.3c0.1,0.4,0.4,0.9,0.3,1.1c-1.4,2-3.4,1.3-5.3,1.1    C173.2,254.3,173.4,254.4,173.4,254.4z"
	pathData30 := "M174.9,332.2c2.4-0.3,4.6-0.7,7.2,0C179.3,332.9,177,332.5,174.9,332.2z"
	pathData31 := "M149.7,162.8c-1-0.4-0.8-1.1-0.7-1.7c0.1-0.7,0.7-1,1.1-0.7c0.3,0.2,0.4,0.9,0.5,1.4    C150.6,162.6,150.1,162.8,149.7,162.8z"
	pathData32 := "M199.9,98.9c3.1-2.1,5.1-6,9.9-5c-1.9-1.4-3.8-2.8-6.3-1.3c-0.4,0.3-0.7,0.7-1.3,0.2c-2-1.8-3.9-1.7-5.9,0.2    c-0.7,0.6-1.8,0.7-2.3,0c-0.6-0.9,0.4-1.5,1-2c1.6-1.1,3.4-1.1,6.5-1c-7.1-4.1-11.6-8.7-9.9-17.1c1.7,5,1.6,10.3,6.3,13.2    c0.8,0.5,1.3,1.8,2.4,1.4c2.8-1.3,4.1-0.2,4.6,2.5c0.2,1.3,1.3,0.9,2,1c2.4,0.3,5.1-0.1,6.1,3.1c0.1,0.2,1.6,0.3,1.9,0    c1.6-1.8,4.5-2.8,3.3-6.5c-0.9,2.4-2.5,3-4.7,3.1c4.4-1.5,3.4-5.7,4.5-9.1c3.7,5.9,3.7,8.9-0.1,13.3c1.1,0.4,1.7,0.5,2.5-0.9    c1.6-2.6,3.2-5.1,4-8c0.2-0.8,0.6-1.7,1.4-1.5c0.7,0.2,1.3,0.9,0.9,2c-1.4,3.7-3.1,7.3-5.9,10.2c-2.1,2.2-5,0.1-7.2,1.4    c-2.3,1.4-4.5-0.2-6.7,0.4C204.8,98.7,202.6,99.9,199.9,98.9z"
	pathData33 := "M192,45.4c1.3-0.9,1.6-3,3.4-1.1c4.2,4.5,10.1,4.4,15.5,5.6c1.6,0.3,3.3,0.1,4.9,1.8c-1.4,0-2.5,0-3.5,0    c-2.3,0-4.3-0.3-5.9,2.7c-2.3,4.4-6.9,4.4-9.8,0.3c-2-2.8-5.3-2.3-7.8-4.6c1.7-0.1,3.3,0.2,4.4-0.4c1.5-0.9,3-0.3,4.5-0.8    C196.2,47.5,194.5,46.2,192,45.4z"
	pathData34 := "M226.9,53.5c3.1-1.5,5.5-3.5,7.8-5.7c2.3-2.2,10.4-1.2,12.2,1.4c0.6,0.9,0.4,1.9-0.8,1.6c-5.9-1.5-10.7,1.7-15.9,3.2    C229.1,54.3,228.1,54.2,226.9,53.5z"
	pathData35 := "M186.6,70.5c4.6,6,1.6,15.7,10.9,18.1c-4.8,0.2-7.2-2.8-9.6-5.7c-1.5-1.9-2.7-4.1-4.2-6.1c-0.6-0.8-1.4-1.5-0.3-2.4    c1.4-1.1,1.5,0.3,1.9,1c0.7,1.1,1.5,2.1,2.3,3.4C188.4,76,185.7,74.1,186.6,70.5z"
	pathData36 := "M173.1,231.7c-6.6,2.5-13.9,1.7-19.1,7.1c-0.5,0.5-1.3,0.8-1.6-0.2c-0.2-0.6-0.3-1.6,0.5-1.8c3.8-1,6.4-5,10.9-4.4    c1.3,0.2,3,0.6,2.4-2.3c-0.4-2,1.7-0.6,2.1-0.1C169.6,231.4,171.1,232.1,173.1,231.7z"
	pathData37 := "M172.8,188.3c0.8,3.7-0.2,7.6,1.5,11.2c0.2,0.5,0.2,1.6-0.8,1.4c-0.6-0.1-1.2-0.7-1.5-1.2    C170.1,195.9,171.1,192.1,172.8,188.3z"
	pathData38 := "M185.6,61.5c4.1-2.9,6.3,1.4,9.6,2.6c-2.4-1.1-3,3.1-5.2,1.7C188.3,64.8,187.1,63.1,185.6,61.5z"
	pathData39 := "M188.6,185.9c0.1,5,3.8,9.5,2.3,15.2C189,195.9,188.8,190.9,188.6,185.9z"
	pathData40 := "M140.9,135.2c-0.2,5-2.8,10-5.3,10.4C137.5,142.2,137.8,137.9,140.9,135.2z"
	pathData41 := "M245.7,225.3c-2,0.7,0.4,4.1-2.4,3.6c-2.4-0.4-0.9-2.8-1-4.2c-0.1-1-1.3-3.2,1.2-3.2C246.1,221.5,243.8,224.5,245.7,225.3    z"
	pathData42 := "M222.6,317.8c-4,1.4-6.4-1-9.1-2.5c-0.6-0.3-1.3-1-0.9-1.7c0.5-0.9,1.7-0.8,2.1-0.3C216.7,315.9,219.9,316.2,222.6,317.8z    "
	pathData43 := "M203,227.2c1.5-3.2,1.3-7.1,4.4-10C207.6,221.6,206.5,224.9,203,227.2z"
	pathData44 := "M142.6,238.9c-3.2,3.1-7.4,4.7-11.1,7c3.2-3.2,6.5-6.2,11.2-6.9L142.6,238.9z"
	pathData45 := "M180.4,198.5c2.6,3.3,2.7,6.7,2.8,10.2C180.7,205.9,180.1,202.6,180.4,198.5z"
	pathData46 := "M193,248.1c-1.1,3.2-2.2,6.4-3.6,10.4C188.7,254.5,189.9,250.6,193,248.1z"
	pathData47 := "M145.9,233.9c3.8-1.5,7.7-2.3,11.8-2C153.9,233.6,150,234.2,145.9,233.9z"
	pathData48 := "M188.8,244.7c1.8-3,2.6-6.6,6.6-9.5C193.2,239.4,192.2,242.8,188.8,244.7z"
	pathData49 := "M176.5,182.3c-1.5-3.3-1.7-6.5,0-10.4C176.5,175.9,176.5,179.1,176.5,182.3z"
	pathData50 := "M246.9,271.4c-1.6-3.9,1.5-5.3,3.6-7.8C250.9,267.4,248.6,269.1,246.9,271.4z"
	pathData51 := "M206,297.6c1.6,2.5,3.8,3.8,6.1,5.4C207.8,303.8,206.3,302.5,206,297.6z"
	pathData52 := "M230.4,190.8c1.9,0.7,1.9,2.7,2.9,4.2c0.6,1-0.5,1.6-0.6,1.5c-1.4-0.8-2.8-1.7-3.1-3.5    C229.5,192.3,229.1,191.3,230.4,190.8z"
	pathData53 := "M176.8,195.1c0.6,2.5,2.2,4.8,1.3,8.1C176,200.5,175.7,198,176.8,195.1z"
	pathData54 := "M150.9,252.3c0.5-3.2,3.5-3.9,5.4-5.6l-0.1-0.1c0.9,4.5-2.8,4.6-5.4,5.6L150.9,252.3z"
	pathData55 := "M194.3,74.6c0.8,2,2.4,3,3.5,4.4c0.5,0.6,0.7,1.7-0.1,2.3c-0.7,0.5-1.5-0.2-1.7-0.8C195.4,78.6,192.8,77.4,194.3,74.6z"
	pathData56 := "M188.9,53.7c1.2,2.6,2.9,3.9,4.4,5.3c0.5,0.5,1.9,0.4,1.2,1.6c-0.5,0.7-2,0.7-2.2,0C191.5,58.4,188.6,57.5,188.9,53.7z"
	pathData57 := "M149.5,228.4c4.4,0,8.9,0,13.3,0c0,0.1,0,0.3,0,0.4c-4.5,0-9,0-13.5,0C149.4,228.7,149.5,228.5,149.5,228.4z"
	pathData58 := "M234.3,263.6c-2.4-2.4-2.5-4.3-1.8-7.4C233.9,258.8,234.7,260.7,234.3,263.6z"
	pathData59 := "M125.4,135.1c-0.6-6,2.9-1.5,5.1-2.1c-0.9,1.1-1.4,1.7-1.9,2.3c-0.2,0.3,0.2,1.3-0.8,0.8c-0.3-0.2-0.8-0.8-0.8-0.9    c0.4-0.8,1.5-1.7-0.3-2c-0.4,0-0.9,1.1-1.4,1.7L125.4,135.1z"
	pathData60 := "M166.7,95c-1.7,0-3.1-0.1-4.4,0c-1,0.1-1.6-0.5-2-1.1c-0.1-0.1,0.3-0.7,0.6-1c0.7-0.6,2.1-0.7,2.2-0.2    C163.3,94.9,165.5,93.5,166.7,95z"
	pathData61 := "M243,237.4c-1.2-0.9-1.6-1.2-1.6-1.3c1-1.6-1.9-3-0.2-4.5c0.2-0.2,1-0.2,1.2,0C243.7,233.3,242.9,235,243,237.4z"
	pathData62 := "M161.4,239c3-0.9,5.8-3,10.3-1.8C167.5,237.6,164.8,240.1,161.4,239C161.3,238.9,161.4,239,161.4,239z"
	pathData63 := "M201,59.4c2.7,0.1,4.4-0.8,6.4-1.1C205.8,60,204.5,62.2,201,59.4z"
	pathData64 := "M225.3,264.7c0.5,1.7,2.3,3,1,5.4C224.7,268.4,223.9,266.8,225.3,264.7z"
	pathData65 := "M142.7,239c2.3-1.3,4.5-3.1,7.9-2c-2.8,1.1-5.1,2.4-8,1.9C142.6,238.9,142.7,239,142.7,239z"
	pathData66 := "M172.6,95.6c-2,1.2-4,2-4.9-1.4C169.5,94.4,170.9,95.5,172.6,95.6z"
	pathData67 := "M237.9,259.6c-0.9-0.4-1.6-1-1.6-2.1c0-1,0.7-1.9,1.2-1.8c1.2,0.2,1.4,1.4,1.3,2.5C238.8,258.8,238.9,259.6,237.9,259.6z"
	pathData68 := "M229.4,262.2c-0.5,1.8,2,3.1,0.4,4.9c-0.3,0.3-0.9-0.1-1.2-0.7C227.9,264.9,227,263.4,229.4,262.2z"
	pathData69 := "M137.3,237.2c2.1-0.8,3.6-3.2,6.5-2.8C142.1,236.6,139.5,236.5,137.3,237.2z"
	pathData70 := "M130.4,254.5c-0.8-3,1.8-3.4,2.8-4.8C133.2,251.7,132.3,253.1,130.4,254.5z"
	pathData71 := "M158.8,175.8c2.2,2.6,0.2,4.3-0.2,6.8C158.1,180,158,178.2,158.8,175.8z"
	pathData72 := "M216,295.2c-0.8-0.1-1.7-0.7-1.6-1.7c0-0.5,0.7-1,1.1-1.6c0.7,0.8,1.7,1.5,2.1,2.4C217.9,295.2,216.9,295.2,216,295.2z"
	pathData73 := "M168,104.9c-0.6,0.1-1.2,0.2-1.8,0.4c-0.9,0.3-1.3,2.1-2.6,1.2c-0.3-0.2-0.4-1.1-0.2-1.5c0.6-1.3,1.9-0.5,2.9-0.7    c0.5-0.1,1.1,0,1.6,0C167.9,104.5,168,104.7,168,104.9z"
	pathData74 := "M150.8,252.1c-2.1,0.8-3.8,3.3-6.5,2c0,0,0.1,0.1,0,0c2-1.4,3.9-3,6.6-1.9C150.9,252.3,150.8,252.1,150.8,252.1z"
	pathData75 := "M156.3,246.7c1.6-1,2.7-3,5.1-2.5l-0.1-0.1c-0.8,2.7-2.7,3.1-5.1,2.5C156.2,246.6,156.3,246.7,156.3,246.7z"
	pathData76 := "M143.7,226.4c2-1.7,3.7-1.9,5.6-1.4C147.9,226.8,146.1,226.5,143.7,226.4z"
	pathData77 := "M161.4,244.1c1.9-0.8,3.1-3.2,6-2c-2.1,0.9-3.6,2.9-6.1,1.9C161.3,244,161.4,244.1,161.4,244.1z"
	pathData78 := "M215.1,298.7c-0.3,1-1,0.8-1.6,0.7c-1-0.2-1.8-0.8-1.6-1.9c0.1-0.5,0.9-0.8,1.2-0.4C213.6,297.9,215,297.4,215.1,298.7z"
	pathData79 := "M138.1,247.1c1.1-2.3,3.2-2.8,5.5-3l-0.1-0.1c-1.4,1.7-3.3,2.7-5.5,3L138.1,247.1z"
	pathData80 := "M187.5,206.4c1.5,1.9,0.6,3.3-1,4.6C185.6,209.2,186.7,207.8,187.5,206.4z"
	pathData81 := "M221.8,272.9c2.2,1.8,2.2,3,0.2,4.6C222,275.9,221.9,274.7,221.8,272.9z"
	pathData82 := "M216.8,134.2c-0.3,0.3-0.7,0.7-1,1.1c-0.5-0.6-1.1-1.2-1.5-1.9c-0.2-0.5,0.1-1.2,0.8-1C216,132.5,216.6,133.1,216.8,134.2    z"
	pathData83 := "M125.9,252.9c0.2,1.4-0.5,1.6-1.4,1.6c-0.6,0-1.1-0.3-1-0.9c0.2-0.8,0.8-1.4,1.7-1.5C126,251.9,125.8,252.7,125.9,252.9z"
	pathData84 := "M156.7,224.2c-2.2,0.9-3.9,0.6-5.6,0C152.8,224.2,154.5,224.2,156.7,224.2z"
	pathData85 := "M138,247c-0.8,0.9-1.7,1.8-2.7,2.8c-0.5-2.9,0.7-3.3,2.8-2.7C138.1,247.1,138,247,138,247z"
	pathData86 := "M220.2,286.3c0,0.6,0,1.2-0.7,1.2c-1,0.1-1-0.7-1-1.4c0-0.6,0.1-1.1,0.9-1C220.1,285.2,220.1,285.7,220.2,286.3z"
	pathData87 := "M245.7,215c-0.2,0.4-0.3,1.1-0.7,1.2c-0.9,0.3-1.2-0.4-1.1-1.2c0-0.6,0.1-1.1,0.7-1.2C245.4,213.7,245.5,214.3,245.7,215z    "
	pathData88 := "M141.6,153.3c-1.9-1-2.8-1.5-3.8-2C139.2,150.8,140.4,151,141.6,153.3z"
	pathData89 := "M141.5,232.2c1.3-2.1,2.1-2.7,3.5-2.2C144.6,231.2,143.5,231.3,141.5,232.2z"
	pathData90 := "M125.3,135c-0.6,0.8-1.2,1.7-1.7,2.5c0-1.3-0.7-3.3,1.9-2.4C125.4,135.1,125.3,135,125.3,135z"
	pathData91 := "M143.6,244.1c1-0.5,2-1.1,3.4-1.8c-0.7,2.5-2.1,2.1-3.5,1.7C143.5,244,143.6,244.1,143.6,244.1z"
	pathData92 := "M150.7,104.4c-0.7,0.8-1.4,1.6-2.2,2.3c-0.2-1.8,0.4-2.7,2.3-2.2C150.8,104.5,150.7,104.4,150.7,104.4z"
	pathData93 := "M112.8,95.1c-1.8,1.8-1.6,4.6-3.7,6.9C109.2,98.7,109.7,96.2,112.8,95.1z"
	pathData94 := "M78.5,83.8c0.1-2.3-2.9-4-1.3-6.7C78.8,79.2,79.2,81.4,78.5,83.8z"
	pathData95 := "M127.2,105.2c0.7-4.5,3.9-1.9,6-2.2C131.3,103.8,129.3,104.5,127.2,105.2z"
	pathData96 := "M124.3,112c-0.2,2.2-0.2,4.4-1.5,6.4C122,115.9,123.2,114,124.3,112z"
	pathData97 := "M107.8,109.7c-0.7-1.8-0.8-3.1,0.4-5.6C109.1,106.6,109.6,108,107.8,109.7z"
	pathData98 := "M124.6,110.9c-1.3-1.8,0.8-2.5,0.8-3.7C126.1,108.6,126.4,109.9,124.6,110.9z"
	pathData99 := "M112.8,99.6c0.3,1.7,1.1,3.5-1.5,4.7C112.1,102.5,112.8,101,112.8,99.6z"
	pathData100 := "M128.6,98.6c-0.2,2.3-0.9,2.9-2.9,3.3C126.8,100.7,127.4,100,128.6,98.6z"
	pathData101 := "M203.1,55.1c-2.7,1-3.7,0.8-4.8-1.4c-0.4-0.9-2.1-2.2-0.3-3.1c1.7-0.9,3.8-1.6,5.6-0.5c1.6,1,1.9,2.5,0.5,4    c-1.1-1-2.3-1.9-4.1-1.4C201,53.7,202.4,53.9,203.1,55.1z"
	pathData102 := "M203.1,55.1c-0.3,0-0.6,0-0.8,0c-2-0.1-3-1.3-3.3-3.1c-0.1-0.4-0.2-1.1,0.8-0.6c1.5,0.8,5.1-1.7,4.3,2.7    C203.8,54.4,203.4,54.8,203.1,55.1z"
	canvas.Translate(0, 10)
	canvas.Scale(0.9)
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
	canvas.Path(pathData99, "fill:"+lineColor)
	canvas.Path(pathData100, "fill:"+lineColor)
	canvas.Path(pathData101, "fill:"+hexcode)
	canvas.Path(pathData102, "fill:"+lineColor)
	canvas.Gend()
	canvas.Gend()
}
