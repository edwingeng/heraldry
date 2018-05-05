package heraldry

import svg "github.com/ajstarks/svgo"

func renderFoxSejantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M42.8,66.4c-2.4,4.6-6.4,6.4-7.3,10.3c6.5,0,7.1-7.4,12.2-9.4c-1.8,2.4-3.7,4.9-5.9,7.8c8.3-3.9,16.8-4.4,24.8-7.1    c10.6-3.6,19.1-9.3,23.4-20.2c1.6-4.2,4.3-7.5,8.3-9.8c3.3-1.9,4.6-1.1,4.6,2.5c0,4.1-0.2,8.2,1.4,12.2    c7.7-3.3,13.5-10.2,22.3-12.6c0.9,9.4,0.6,18.6-2.9,27.3c-1.2,3.1-1,5.7,0.5,8.4c3.4,6.1,5,12.5,4.7,19.6    c-0.4,9.2,0,18.3-0.1,27.5c-0.1,3.4,0.5,6.6,2.1,9.3c6.2,10.2,10.9,21.1,14.8,32.3c5.1,14.6,13.8,27.5,20.6,41.4    c4,8.2,8.4,16.2,12.3,24.4c6.7,14.2,11.8,29,14.5,44.6c1.4,7.9,1.9,17,10.7,21.7c9.4,5,21.1,1.8,26.2-9.1c4-8.7,2.4-17.1,1.1-25.8    c-2.1-14.2-8.4-27.2-11.6-41.2c-2.9-13-2.5-25.6,1.5-38.1c1.9-6.1,3-12.4,6.6-18c4.5-6.9,11.1-11.6,17.6-16.1    c2.9-2.1,4.2-5.4,8.4-8.2c2.5,23.7,10.3,46,5.8,69.9c3.8-1.1,3.8-1.1,7.5-0.4c-4.6,6.4-3.4,13.4-1.7,20.1    c2.4,9.5,3.5,19.1,4.5,28.7c0.3,2.9-1.4,5.5-2.1,8.2c2.5,0.6,2.8-1.9,6.4-2.8c-5.8,7.6-3.6,15.9-6,23.3c1.8,0.7,2.7-1.2,4.6-2.2    c0.2,4.8-2.2,8.5-4.1,12c-8.9,16.3-22,27.8-40.1,32.9c-8.2,2.3-16.4,1.5-23.7-2.2c-7.9-4.1-16.2-3.2-24.1-2.5    c-6.3,0.6-12.7,0.9-19.1,1.4c-11.9,0.9-24,0.6-36,0.2c-10.5-0.3-21,0.1-31.5-1.8c-3-0.5-2.6-2-2.1-3.7c1.7-5.8,6.9-9.8,13.1-9.4    c6.3,0.5,12-2.8,18.7-2.1c-3.5-5.3-9.1-7.6-12.2-12.6c-8.2-12.7-8.8-34.8-0.5-47c4.2-6.2,10.6-8.8,18.3-9.8    c-2.1-2.2-4.5-2.4-6.4-2.4c-5.6,0-11.1-0.2-16.2-2.7c-2.7-1.3-2.5,1.5-2.7,2.2c-2.6,7.7-5.3,15.5-7.1,23.4    c-1.2,5.3-2.8,10.6-4.3,15.8c-3.5,11.8-3.6,24.6-9.9,35.6c-3.5,5.9-6.4,13.2-14.2,14.2c-7.9,1-16,0.6-23.9,0.2    c-3.3-0.1-5.3-4.7-3.3-7.6c2.2-3.3,4.5-6.8,8.5-8.6c3.1-1.4,5.3-3.8,7.1-6.6c5-8.1,7-16.9,8.3-26.4c2.6-19,1.4-38,1.6-57    c0.1-11.6-3.1-22.9-4.4-34.3c-0.5-4.5-0.8-8.7,0.5-13.1c3.4-11.3,5.5-23.1,11.7-33.5c3-5.1,5.9-10.2,7.4-16.1    c2.9-10.9-4.4-26.5-18.9-27.8c-11.8-1.1-24.2-0.3-34.9-7C26.5,87.6,24,87,24,84c0-3,2.4-4,4.4-4.7C34.7,77.3,36.5,70.4,42.8,66.4z    "
	pathData1 := "M260.6,249.6c2.2,0.8,2.3-2.1,4.7-2.3c-1.8,3.8-2.2,7.7-1.9,11.7c0,0.4-0.5,1.1-0.2,1.7    c1.8,3.3,0.6,8.5-2.6,11.1c-5.3,4.3-5.3,4.3-7.4,11.4l-0.2-0.1l0-0.1c4.2-3.1,6.6-8.1,11.1-10.9c1.1-0.7,1.6-2.1,4.1-2.4    c-2.9,3.1-5.4,5.7-6.2,8.9c-1,3.9-6.8,2.5-7.1,6.9c1.4,1,1.9-0.2,2.6-0.8c0.5-0.4,0.9-1,1.7-0.6c2.5,1.2,2.7,1.2,4.6-2.5    c0.6-1.1,1-2.2,1.9-3.2c-0.8,6.5-2.8,12.4-6.6,18.1c-2.8,4.2-9.1,5.1-10,10.9c-1.8-0.3-3,1-4.4,1.8c1.4-0.8,2.6-1.8,4.3-1.9    c7.3-3.2,12.6-8.9,17.3-15.1c-4,10.8-16.8,25.2-24.1,27.7c-0.2-1.5,2.2-1.6,1.2-3.6c-4,0.6-7,3.8-10.7,5.2    c-3.9,1.5-8.2,0.2-12.9,0.8c4,5.7,9.1,0.9,13.4,2.3c-5.4,1.5-10.1,4.8-16.5,2.7c-6.5-2.2-13.2-3.9-19.4-7    c-9.7-5-12.5-10.9-9.8-21.5c1.2-4.6-1.4-12.1-5.2-14.8c-1,2.3,0.4,5.4,0.9,6.3c2.9,5,0.5,9.3-0.2,14c-0.6,4.4,1.7,8.4,4.4,11.9    c1.2,1.6,2.9,2.8,4.1,4.6c-3.6-1.3-6.6-3.4-8.7-6.5c-0.4-0.5,0.5-1.9-0.8-1.7c-0.8,0.2-1.5,1-1,2.3c1,2.3,2.7,3.8,6.2,5.7    c-5.6,0.6-9.7,1.1-13.8,1.5c-4.5,0.4-9.1-0.9-13.5,1.8c-2.2,1.4-5.7,0.3-7.8-0.2c-6.5-1.6-12.7,0.6-19,0.5    c-3.7-0.1-5.1-1.1-3.7-5.4c-4.5,5.4-10.7,3.7-16.1,5.5c-0.4-4.7,4.9-4.4,6.1-8c-6.4-0.4-8.3,9.3-16.1,6.2    c6.5-4.1,10.9-10.5,20-8.6c5,1,10.3,1.1,15.4,2.2c4.4,0.9,8.9,0.1,13.2-1.1c6.2-1.8,12.3-3.3,18.7-2.8c1.1,0.1,1.9-0.2,2.8-1.4    c-1.1-1-3.1-0.6-3.8-2.1c-0.2-0.4-0.9-1.1-1-1c-7,3.8-6.8-4.6-10.6-6.3c-0.6,1.7-0.6,3.2,0.8,4.5c1.1,1,3.1,2.5,1.3,3.9    c-1.4,1.1-3.7,0.9-5.3-1c-1.4-1.7-1.5-1.6-2.6,0.8c-0.4,1-3.6,0.8-5.1-0.5c-2.6-2.3-4.2-5.4-6.2-8.1c-0.5-0.6-1-2.6-2.5-1    c-1.3,1.3,0.4,1.6,0.8,2.2c1.7,2.6,3.5,5.1,6,8.8c-8.4-4.3-10.7-11.7-14.1-17.5c-0.9,4.6,0.9,8.8,8,17    c-6.6-2.8-10.6-6.5-10.7-14.1c-2.1,3.7-1.7,5.7,1.3,11.8c-5.9-3.1-6.1-9.3-7.7-14.8c1.8-0.8,2,1.6,3.4,1.3c0-1.9,0-1.9-5.7-5.4    c0.3,4.8,1.4,9.1,3.5,14.5c-7.9-7.2-5.9-14.2-2.2-21.4c-5.8,4-6.9,7.9-4.3,15.7c-1.2-0.9-2.3-1.4-2.7-2.2    c-6.1-13.2-9-26.5-2.2-40.4c2.7-5.5,14.6-11.3,23-9.2c4.5,1.1,9,2.3,13.4,3.6c8.2,2.3,10.2,10.3,15.7,15c1.7,1.4-0.8,3.7,1.1,5.1    c0.5,0.4,1,1,1.6,1.1c0.8,0.1,1-0.5,1.1-1.3c0.2-2.7,0.4-5.4,0.6-8.1c3.1,0.4,2.4,3.4,3.4,5.1c-1-1.8-0.1-4.7-3.4-5.1    c-0.8-0.4-0.3-1.2-0.6-1.7c-1-1.5-1.1-4.4-3.7-3.7c-1.8,0.4,1,2.7-0.8,3.7c-2.4-1.2-1.6-3.8-2.2-5.7c-1.2-3.8-0.5-8.1-7-7.8    c-1.9,0.1-3.9-4.2-5.9-6.4c-0.6-0.7-0.2-2.1-2.2-2.3c-1.3,0.3,0.7,2.5-1.1,3.3c-0.7-3.2-3.8-5-4.4-8.3c-2.4,2.8-3.3,2.1-6.5-5.2    c-0.9,5,0.6,9,4.1,12.6c-5.1-1.7-8.8-5.4-12.4-9.3c-2.3,2.6,2,3.1,0.8,5c-2.4-0.3-3.7-2.2-5.3-3.7c-1,1.2-1.2-5.6-2.8-0.3    c-0.2,0.7-1.5,0.4-1.6-0.7c-0.1-2.5-2.7-3.5-3.5-6c0.8,3.7-2.7,3.2-3.9,4.8c-0.3,0.4-1.7,0.4-1.5-0.4c1.1-4.6-5.9-9.9,1-14.1    c0.6-0.4,0.2-0.7-0.2-1.3c-2.2-3.4-3.3-7.2-3-11.4c-1.8,1.5-1.8,1.5-2.4,3.3c-1.1-1.5,0.4-4-2.3-4.5c-1.4,4.7,2.6,9.7-0.8,14.3    c-0.6-1.1-1.2-2.2-1.8-3.3c-3.1,2.1-0.9,3.9,0.7,5.8c-1.6,0.9-3-1.3-4.1,0.3c0.9,2.7,4.3,2.4,5.8,5.1c3,5.3,0.5,9.9-0.8,14.6    c-4.1,14.9-8.9,29.5-12,44.6c-1.4,6.8-3.6,13.4-4.8,20.2c-1.3,7.2-5.8,13.9-11.8,18.4c-4.8,3.6-11.3,2.7-17.6,2.5    c3.1-4.8,6.9-8,13.1-7.7c-6.6-1.9-10.9,0.9-14.5,5.5c-0.8,1-2,1.6-3.2,2.2c-0.9,0.4-1.5-0.3-2.1-1c-0.7-0.9,0.3-1.2,0.7-1.5    c2.8-2.1,4.3-5.6,8.3-6.6c9.4-2.4,12.4-10.7,15.5-18.4c2.9-7.3,3.5-15.1,4.1-22.8c0.7-7.9,1.4-16,3.1-23.7    c1.9-8.4,1-17.1,3.3-25.6c0.5,0.8,0.8,1.6,1.3,2c1.1,0.7,1.6,3.2,3.1,2c1.5-1.1,0.5-2.7-0.9-3.8c-0.6-0.5-1.4-1.1-0.3-2.6    c2.7,4.9,9,6.7,9.9,13c0.1,0.7,1.4,1.2,2.2,1.9c0.4-0.7,0.8-1.3,1.1-2c1.1-2.6,2.6-5.3-0.5-7.6c-0.1-0.1-0.3-0.1-0.5-0.1    c3-6.6-6.8-4.8-6.1-9.9c-0.8,3.7,1.3,6.4,2.6,10.1c-4-1.3-5.7-4-7.1-6.7c-0.6-1-0.8-2.3-2.1-2.4c-2-0.2,1.5,3-1.6,2.3    c-4.8-11.5-10.8-23.2-2.8-36.2c-1.1,0.9-2.2,1.8-3.5,2.9c0-2.1,1.7-3.7,0.6-5.7c-1-1.2-1.1,5.3-2.8,0.5c-0.6-1.7-1.3,0-1.6,0.3    c-0.7,0.7,0.1,1.4,0.4,2.1c2.3,7.6,1.5,14.8-2.3,21.8c-1.5,2.8,1.1,4,2.5,5.2c3.6,3,5.6,7.3,4.8,11.3c-0.9,4.4,0.5,8.7-1.2,13    c-2-1.1-2.6-3.5-3.9-5.4c-0.5-0.7-1.2-1.5-2-1.1c-1,0.6-0.9,1.8-0.3,2.6c1,1.3,2.2,2.5,3.5,3.5c2.2,1.7,2.7,3.8,1.7,6.9    c-1.9-2.7-3.8-5-4.9-8.1c-2.5,4,0.8,6.5,2.6,7.8c4.7,3.5,2,7.7,1.6,11.3c-0.8,7.8-2.4,15.4-2.9,23.2c-0.6,11.2-3.2,22-10.2,31.2    c-0.8,1.1-1.7,2.4-2.8,2.8c-6.3,2.3-11.8,5.6-15.7,11.3c3.1-3.5,4.2-8.3,9.4-10.2c6.8-2.5,8.7-9.5,11.2-15.6c3-7.1,3-14.9,4-22.4    c2-14,0.7-28.2,1-42.2c0-2.3,2-4.5-0.3-7c3.6,1,4.7,4,7.3,6c-0.6-5.4-8.2-7.6-5.6-13.8c1.3,0.1,1.6,1.7,2.6,1.9    c1.2,0.3,0.3,3.3,2.2,1.9c2.1-1.5,0-2.6-1-3.4c-6.8-5.7-6.5-13.6-7.3-21.6c1.4,0.8,1.1,1.7,1.1,2.5c-0.1,1.8,0.8,3.1,1.8,4.3    c0.1,0.1,0.9-0.2,1.2-0.5c1.6-1.5,1.3-3.4,0.6-5c-1.1-2.5-2.7-4.8-0.9-7.8c0.5-0.8,1.9-2.2,0-3.3c-0.7,1.7-3.5-1-3.7,1.8    c0,0.8,0.2,1.7-1.4,1.8c-2-4.2,2-6.4,3.9-9.3c-2.9-1.4-2.9,1.2-4.8,2.3c1.4-6.7,1.7-13.1,6.4-18c1.4-1.5,1.8-3.3,2.9-4.9    c0.3-0.5,0.3-1.8,0.4-1.8c6-0.3,5.6-5.5,7.4-9.8c-5.2,3.5-8.6,7.9-11.5,13c-0.4-0.3-0.7-0.5-1.1-0.8c0.9-1.7,1.5-3.6,2.6-5.1    c2.7-3.6,1.8-8.3,4-12.2c1.8-3.3,6.2-2.9,7.2-6.7c-1.7-0.8-2,2.2-3.7,1.1c0.6-3.9,7.3-3.4,6.4-8.6c-0.9,0.6-1.7,1.1-3.5,2.3    c4.6-9.8,2.7-18.4-1.2-27.1c4.1,3.3,7.4,8.1,13.7,6.4c-1.1-2.5-4.4-1.9-5.8-4.5c3.7,0.3,6.2,3.2,9.9,3.9c-1.1-3.1-5.4-2.1-5.9-5.6    c3.1,0.3,5.2,3.6,8.7,2.3c-0.8-1.2-3-0.2-3.4-2.4c4-2.6,9.7,1.5,13.8-3.2c-3.5,0.8-6.4,1.2-9.6,0.1c-1.5-0.5-4-1.7-5.6,1.2    c-0.9,1.7-2.6,0.2-3.9-0.3c-1.8-0.8-3.8-1.4-5.7-2.1c-0.1,2.2,2.7,1.5,2.6,3.4c-1.1,1.1-2.8,1.1-3.5,0.3    c-3.8-3.8-10.4-2.6-13.3-7.9c-1.1-2.1-2.8,1-4.4,0.3c-6.4-3-13.9-0.5-20.7-5c7.8,0,14.7,0.7,21.3,3.9c-1.3-5-7.8-9.5-13-9    c-0.6,0.1-1.6,0.1-0.4,0.9c2,1.4-0.3,2-0.6,1.9c-5.6-0.9-11.9,2.5-17-2.5c0.9,1.1,1.8,2.1,2.9,3.4c-2.5,0.2-3.7-2.1-5.9-2.4    c-2-0.3-2.5-3-2.8-4.8c-0.3-2,2-1.4,3.2-1.9C47.8,75.1,59.8,73.4,71,69c5.7-2.2,11.2-4.9,16.7-7.3c2.9-1.2,5.7-2.3,9-1.4    c1,2.4-2.7,1.5-2.2,4.4c10.7-5.3,17.3-16.3,29.3-19.5c0,6.2,2,12.7-5.8,17c2.5,0.8,3-1.8,5.3-0.6c-2.7,3.7-6.8,5.4-10,8.1    c2,1.6,2.9-1.1,4.5-1.4c0.1,4.4-6.5,3.8-5.7,8.4c2.5-0.2,3.8-2.2,5.6-3.7c0.5,3.9-2.7,4.7-4.7,6.6c1.3,0.5,2.8,1.5,3.6,1.1    c5.7-3.1,6,1.6,7.5,4.9c3.3,7.4,1.2,15.3,2,23.1c-2.8-1.4-1.6-4.9-4-6.3c-0.5,3.4,0,6.4,1.9,9.3c1.2,1.8,1.1,1.9-0.7,4.7    c-0.3,0.5-0.4,1.3-0.3,1.9c1.9,8.6,3.8,17.2,9.2,24.4c4.6,6.3,7,13.7,9.7,20.9c3.9,10.5,9.3,20.4,14.4,30.4    c11.4,22.5,22.7,45.1,29.8,69.5c2.5,8.6,3.8,17.5,6.6,26.1c2.4,7.3,9.4,12.4,16.7,12.5c1.3,0,3.1-0.8,3.9,1.2c-0.4,0-0.5,0-0.5,0    c0.1-0.1,0.2-0.2,0.4-0.2c1.8-2.2,4.7-1.5,6.9-1.9c8.3-1.6,16-14.5,16.1-21.8c0-4,1.7-7.9,0.7-12c-1.5,0.1,0.1,2.1-1.8,2    c0.5-4.4,0.7-9-0.9-13.4c-0.2-0.6-1.6-3.4,1.8-2.3c0.5,0.2,0.4-2.4-0.9-2.6c-3.1-0.6-2.9-3.3-3.9-5.3c1.3-0.8,1.6,1.1,2.9,1    c-0.8-4.1-4.6-6.4-5.7-10.3c3.1,3.2,5.8,6.7,9.9,9.2c-10.7-12.4-16.6-26.6-16.1-43c2.6-0.3,3.2,1.7,3.9,3.4c0.7,1.7,1.1,3.4,2,6.2    c0.7-5.5-1.2-10-4.2-11.9c-1.9-1.2-1.9-1.2-0.4-6c1.4,1.8,0.9,5.1,4.8,5.7c-1.5-3.9-0.6-8.1-4.3-10.7c-1-0.7-0.2-1.4,0.1-2.1    c1.6-5,3.2-10,4.9-15.2c2.8,1.7,1.6,5.6,5,6.6c0.2-1.2,0.3-2.2-0.1-3.5c-0.9-2.9-3.4-5.6-0.8-9.2c2-2.7,2.7-6,6.1-7.8    c2.3-1.3,3.9-3.9,6.4-5.2c2.3-1.2,4.2-2.9,6.6-5.8c0.2,11.9-6.5,19.4-11.9,27.5c5-4.2,9-8.9,12.6-14.5c0.1,0.6,0.4,1.2,0.2,1.6    c-0.7,1.7,0.4,1.6,0.7,3.6c0.4,3.3,2.1,6.6,1.1,10.1c-1.5,5.2-3.3,10.1-7.6,13.7c-0.7,0.5-1.3,1.1-0.7,2.1c-4.2-1-0.6-3.2-0.7-4.9    c-1.2,2.6-2.7,2.4-4.1,0.2c-0.7-1.1-0.2-2.2,0.1-3.3c1.2-4,3.8-7.2,5.3-10.1c-1.4,2.5-3.5,5.4-4.9,8.7c-1,2.4-1.8,4.5,0.5,6.8    c2,2,0.3,4.3-0.3,6.9c0.9-2.1,1.2-4.6,4-4.4c6.6-1.8,7.5-8.1,10.3-13.6c1.4,11.3-0.7,20.8-9,28.4c4.2-2,7.5-4.5,8.6-9.6    c0.5,3.3,0.2,3.2-0.7,5.3c-3.1,7.3-9.4,13.6-7.3,23.2c1.7-7.5,5.4-13.5,8-20c1.8,10.1-7.7,18-5,28.6c0.9-8.2,7.1-13.6,9.3-21.1    c-1.8,9.9,0.2,19.1,2.9,28.2c0.8,2.5-1.7,2.8-2,4.6c-0.2,1.3-2.5,2.2-0.9,4.1c-0.2,2-0.4,3.9-1.4,5.7    C260.4,253.6,260.7,251.6,260.6,249.6z"
	pathData2 := "M88.5,57.9c4.1-5.6,4.6-13.5,11.7-17.9c-0.8,4.1,1.5,7.9-1.7,11.2c-0.7,0.7-0.4,2,0.4,1.7    c3.1-1.2,2.4,1.1,2,2.3c-0.3,1.2-0.7,2.7-3.1,2C94.9,56.4,91.7,56.7,88.5,57.9z"
	pathData3 := "M95.1,321.7c4.2-4.6,7.2-9.1,15-7.1C103.5,315.1,102.4,322.8,95.1,321.7z"
	pathData4 := "M121.8,56.6c-5.1,0.6-3.4,7-7.3,9.6c-1.3,0.9-6.4,0.1-2.9,4.6c0.4,0.4-3.4,2.1-4.3,1.6c-2.3-1.5-5-4-3.2-6.9    c3.3-5.3,6.4-10.8,12.6-13.5c2.3-1,3.4,0.9,2.9,1.6C115.5,59.5,122.4,53.9,121.8,56.6z"
	pathData5 := "M78.5,77.6c-2.6-0.6-4.6-1.9-7.1-4.9c9.1,1,17.2,1,25.9-0.5c-2.4,1.7-3.2,3.8-4.4,5.7c-2.7,4.4-8.6,4-12.9-0.8    c-0.7-0.8-1-1.2-1.8-0.7C78.1,76.4,78.3,77.2,78.5,77.6L78.5,77.6z"
	pathData6 := "M197.3,307.2c4.2,3.2,5.8,8,10.2,9.8c0.6,0.2,1.6,0.9,1.6,0.9c2.8-4.5,5.5,0.2,8.2,0.1c1.4-0.1,4.2,2.3,3.4-1.9    c-0.2-0.8,0.3-1.1,1.4-0.4c2.3,1.4,4.4,1.6,6.7,0c2.9-2.1,5.7-3.2,8.7,0c-5.8,1.6-11.8,3.2-17.6,5c-1.5,0.5-2.4,0-3.6-0.6    c-1.7-0.8-3.4-1.7-5.5-0.5c-2.2,1.3-4-0.5-5.5-1.9C202.3,314.9,198.7,312.6,197.3,307.2z"
	pathData7 := "M196.2,316.8c0.6,0.1,1.8-0.2,1.6-0.4c-2.8-2.2-4-5.7-6.3-8.5c-1.1,1.8,1.8,3.2-0.5,5.2c-0.5-4.7-3.6-8.4-2-13.2    c0.3-0.9,0.8-1.7,1.5-1.7c0.5,0,1.4,1,1.4,1.7c0.5,10.2,6.4,16.5,15.2,20.6c0.2,0.1,0.4,0.5,0.5,0.8c-3.7,0.5-9.5-1.9-11.3-4.8    L196.2,316.8z"
	pathData8 := "M246.4,191.4c-2.5,1.3-2.9,4-4.1,7c-0.8-4.4,1.2-8.5-2.1-11.9c-1.2-1.2,1-3.5,1.8-5.2c1.5-3.6,4.4-8.1,7.9-11.7    c-1.5,5.5-6.3,8.7-7,13.9c-0.1,0.6-0.7,1.1-0.8,1.7c-0.3,1.4,0.6,2.4,1.7,2.6c1.6,0.3,1.2-1.2,1.2-2.1c0-1.1,0.3-2,1.7-2.5    c1.1,2.1-0.6,3.9-1.2,5.5C244.9,190.4,245.4,190.8,246.4,191.4C246.4,191.5,246.4,191.4,246.4,191.4z"
	pathData9 := "M77.1,304.1c-0.2-3.4,2.4-6,2.5-7.7c0.2-3.6,3.3-4.3,4.3-6.7c0.3-0.6,1-0.9,1.7-0.5c0.3,0.2,0.5,1,0.4,1.2    C82.6,294.6,82.1,300.4,77.1,304.1z"
	pathData10 := "M123.4,222.7c-8.8-0.7-13.1-6.8-9.6-13.8c0.4,3.9-0.7,7.9,4,9.7c1-0.2-0.5-2.1,1-2.4C120.3,218.2,120.5,221.1,123.4,222.7    z"
	pathData11 := "M248.7,198.7c0.7,1.8-0.2,3-1.2,4.4c-1.1,1.6-2.5,3.3-0.4,5.6c2.3,2.5-2.9,4-1.9,5.5c-0.1-4.6-2.5-10.9,3.6-15.4    L248.7,198.7z"
	pathData12 := "M209.6,307.5c2.9-3.5,7.2-1.7,11-2.4c0.7-0.1,1.3,0.1,1.2,0.9c0,0.3-0.7,1-0.8,0.9C217,305.5,213.4,311.5,209.6,307.5z"
	pathData13 := "M239.6,314.7c4.5-0.7,7.5-4.9,12.8-5.2C249.2,314.9,244.1,314.1,239.6,314.7z"
	pathData14 := "M251,273.9c2.8-3.5,4.7-6.9,6-10.7c0.2-0.7,0.3-1.3,1.2-1.1c0.3,0,0.6,0.3,0.7,0.5C260,265,254.7,273,251,273.9z"
	pathData15 := "M257.5,227.6c1.8,4.7,3,9.7,2.2,14.9C258.9,237.5,256.6,232.9,257.5,227.6z"
	pathData16 := "M168.6,261.5c2.6-1.4,3.1,0.6,3.7,2.3c0.7,2,1.9,4.2,1,6.2c-0.9,1.8-1.1-0.2-1.5-0.8C169.9,267,171.2,263.4,168.6,261.5    C168.5,261.5,168.6,261.5,168.6,261.5z"
	pathData17 := "M131.8,222.4c0,2.4,0,4.8,0,8.5c-0.9-4.1-5.5-3.8-5.3-7.8c1.8-0.1,2.4,1.2,2.7,2.8C130.5,224.7,129.5,222.6,131.8,222.4z"
	pathData18 := "M248.8,252.5c1.9-3.4,3.7-6.7,5.5-10.1C254.8,246.1,251.7,252,248.8,252.5z"
	pathData19 := "M73.8,168.8c1.8-3.7,4.8-5.4,8.9-5.5c-3.3,1.3-4.9,5.5-9,5.3L73.8,168.8z"
	pathData20 := "M260.6,249.6c1.3,2.6-0.1,5.1-0.7,7.6c-0.2,1-1.1,2.4-2.5,1.5c-0.9-0.6-0.7-2,0-2.6c2-1.7,2.1-4.3,3.2-6.4    C260.7,249.7,260.6,249.6,260.6,249.6z"
	pathData21 := "M78.4,77.7c0.2,0.9,1.9,1.1,1,2.4c-0.6,0.8-1.3,1.3-2.5,0.5c-0.9-0.6-2.1-0.7-3.1-1.2c-0.9-0.4-2.8,2.1-2.6-0.5    c0.1-1.3,2.6-1.8,3.9-1.1C76.5,78.6,77.4,79.8,78.4,77.7C78.5,77.6,78.4,77.7,78.4,77.7z"
	pathData22 := "M87.7,112.3c3.8,1.4,4,3.4,0.5,7.5C88,117.2,87.8,114.7,87.7,112.3z"
	pathData23 := "M202.8,304.2c-0.5,0-1.4,0-1.4,0c-0.2,4.7-1.6,1.2-2.4,0.3c-0.8-0.8-2-1.3-2-3.4c1.9,1.6,5.4-1.3,5.7,3.2L202.8,304.2z"
	pathData24 := "M73.6,168.6c0.9,3.5,0.9,6.6-3.5,8c0.3-3.1,2-5.5,3.6-7.9C73.8,168.8,73.6,168.6,73.6,168.6z"
	pathData25 := "M116.2,279.5c1.1-2.7-4.3-5.5,0.6-8.2C117,274.1,118.7,276.8,116.2,279.5z"
	pathData26 := "M254.1,235.1c0.8,2.9,0.4,4.1-2.5,8.5C251.4,240.7,251.9,239.1,254.1,235.1z"
	pathData27 := "M93.9,83.5c-2.3,2.6-5.3,1.1-7.9,1.4c-0.8,0.1-1.2-0.5-0.9-1.2c0.1-0.3,0.9-0.7,1.1-0.6C88.6,85.1,91.4,81.9,93.9,83.5z"
	pathData28 := "M144.3,301.4c4.6,2.2,5.4,3,6.6,5.7C147.6,306.1,146.4,303.2,144.3,301.4z"
	pathData29 := "M95.7,242c-1.8-3.2-3.3-5.7-4.8-8.3C92.5,236.4,97.1,237.3,95.7,242z"
	pathData30 := "M71.1,242.3c2.4,1.9,3.4,4.4,4.3,8.1C71.9,247.9,72.3,244.7,71.1,242.3z"
	pathData31 := "M90.8,254.8c0.1-2.4,0.2-4.9,0.3-7.4c2.7,4,2.6,5.3-0.5,7.3L90.8,254.8z"
	pathData32 := "M144.2,183.9c-1.4-1.9-1.5-3.9-0.4-6.7C145.7,179.9,144.7,181.9,144.2,183.9z"
	pathData33 := "M109.9,218.2c1.7,2.6,0.8,5.4,0.7,8.3C108,224,111.7,220.7,109.9,218.2z"
	pathData34 := "M95,249.4c-2.2-3.1-3.6-5.3-4-8.9C93.2,243.4,94.9,245.5,95,249.4z"
	pathData35 := "M155.4,198.6c0.9,2.5,0.9,4.1-0.9,5.2c-0.3,0.2-1.2,0-1.2-0.2C152.9,201.9,154.2,200.8,155.4,198.6z"
	pathData36 := "M249.9,301.1c0.1,2.1-1.2,2.9-2.6,3.6c-0.1,0.1-1-0.8-1-0.9C246.9,302.2,247.8,301.1,249.9,301.1z"
	pathData37 := "M249.2,307.2c-2.2,0.9-3.6,3.8-7.3,2.2c3.1-0.1,4.4-3.6,7.5-2.1C249.4,307.3,249.2,307.2,249.2,307.2z"
	pathData38 := "M202.7,304.3c2.5-1.1,3.9,0.4,5.2,2.3c-2.7,1.3-3.6-1.1-5-2.4C202.8,304.2,202.7,304.3,202.7,304.3z"
	pathData39 := "M149.5,195.5c-0.4-2-1.3-4,1-5.3C152.1,192.4,149,193.6,149.5,195.5z"
	pathData40 := "M146.5,188.5c-0.3-1.5-1.3-3.2,0.3-4.4c0.1-0.1,0.7,0,0.8,0.2C149,186.1,146.8,187.1,146.5,188.5z"
	pathData41 := "M159.6,209.2c1.3,2.2-0.9,3-0.7,5.9C157.7,211.9,157.6,210.3,159.6,209.2z"
	pathData42 := "M90.7,254.7c0.1,0.1,0.2,0.3,0.2,0.4c-0.4,1.1,1.8,2,0.3,3.3c-0.7,0.6-1.1,0.2-1.5-0.3c-1.7-1.8,0.3-2.4,1.2-3.3    C90.8,254.8,90.7,254.7,90.7,254.7z"
	pathData43 := "M120.8,145.5c2.1-0.7,3.7-0.8,5.1,0c0.1,0.1-0.2,0.8-0.3,1.3C124.2,146.4,122.7,146,120.8,145.5z"
	pathData44 := "M141.9,177.1c-1.2-1.8-2.4-3-1.6-4.9C142.7,173.2,141.9,175,141.9,177.1z"
	pathData45 := "M114.1,97.4c-2.5,0.9-4.5,0.6-6.5,0.1C109.6,96.9,111.6,96.8,114.1,97.4z"
	pathData46 := "M213.2,303.1c-1.1,1.5-2.2,3-4.1-0.1c1.7,0.1,3,0.1,4.3,0.2C213.3,303.3,213.2,303.1,213.2,303.1z"
	pathData47 := "M219.1,303.7c0.9-1.7,2.1-2,3.7-1.3C222,304,220.8,304.4,219.1,303.7z"
	pathData48 := "M165,220c0.9,0.1,0.9,0.7,1,1.4c0,0.9-0.2,1.7-1.3,1.5c-0.7-0.1-0.8-0.9-0.7-1.6C164.1,220.7,164.2,220.1,165,220z"
	pathData49 := "M157.4,205.1c-0.4,1,0.1,2.5-1.3,2.8c-0.3,0.1-1-0.4-1-0.7C155,205.7,156.2,205.5,157.4,205.1z"
	pathData50 := "M196.3,316.7c-2.3,1-3-1.1-4.5-2.4c2.6-0.5,3.1,1.6,4.3,2.6C196.2,316.8,196.3,316.7,196.3,316.7z"
	pathData51 := "M254.2,298.2c-0.8,0.9-1.7,1.8-2.5,2.7c-0.2-2.2,0.5-3.3,2.7-2.6C254.3,298.3,254.2,298.2,254.2,298.2z"
	pathData52 := "M254.3,298.3c-0.7-2,0.3-2.9,2.8-2.6c-1,0.8-2,1.6-3,2.4C254.2,298.2,254.3,298.3,254.3,298.3z"
	pathData53 := "M161.6,219.2c-1.2-2.3-0.3-3.2,1-4.1C163.5,216.8,161.8,217.4,161.6,219.2z"
	pathData54 := "M248.9,198.9c0.8-0.9,1.7-1.8,2.5-2.8c0.3,2.2-0.5,3.2-2.7,2.6C248.7,198.7,248.9,198.9,248.9,198.9z"
	pathData55 := "M81.9,184.2c1-1.4,0.7-2.9,3.1-3.3C84.1,182.2,84.7,183.8,81.9,184.2z"
	pathData56 := "M253.1,283.1c-0.8,0.9-1.7,1.8-2.5,2.8c-0.2-2.2,0.5-3.2,2.7-2.6C253.3,283.3,253.1,283.1,253.1,283.1z"
	pathData57 := "M82.2,175.1c-0.2-2.3,0.7-2.9,2.2-3.1C84.8,173.7,83.3,174,82.2,175.1z"
	pathData58 := "M129.5,151c-1.1-0.1-2.2-0.2-3.2-0.3c0.2-0.4,0.4-0.9,0.7-1.2C128.3,148.4,128.4,151.2,129.5,151z"
	pathData59 := "M97.8,139.6c0.5-2.7,1.1-3.7,3.3-3.5C100,137.3,99.3,138.1,97.8,139.6z"
	pathData60 := "M87.5,73.8c1.6,0.5,2.5,1.1,2.4,2.2c0,1.3-1,2-2.3,2c-0.4,0-1.6-0.5-1-0.6C88.9,76.9,87.3,75,87.5,73.8z"
	canvas.Translate(30, 10)
	canvas.Scale(0.8)
	canvas.Path(pathData0, "fill:"+lineColor)
	canvas.Path(pathData1, "fill:"+hexcode)
	canvas.Path(pathData2, "fill:"+hexcode)
	canvas.Path(pathData3, "fill:"+hexcode)
	canvas.Path(pathData4, "fill:"+lineColor)
	canvas.Path(pathData5, "fill:"+lineColor)
	canvas.Path(pathData6, "fill:"+lineColor)
	canvas.Path(pathData7, "fill:"+lineColor)
	canvas.Path(pathData8, "fill:"+lineColor)
	canvas.Path(pathData9, "fill:"+lineColor)
	canvas.Path(pathData10, "fill:"+lineColor)
	canvas.Path(pathData11, "fill:"+lineColor)
	canvas.Path(pathData12, "fill:"+lineColor)
	canvas.Path(pathData13, "fill:"+lineColor)
	canvas.Path(pathData14, "fill:"+lineColor)
	canvas.Path(pathData15, "fill:"+lineColor)
	canvas.Path(pathData16, "fill:"+lineColor)
	canvas.Path(pathData17, "fill:"+lineColor)
	canvas.Path(pathData18, "fill:"+lineColor)
	canvas.Path(pathData19, "fill:"+lineColor)
	canvas.Path(pathData20, "fill:"+lineColor)
	canvas.Path(pathData21, "fill:"+lineColor)
	canvas.Path(pathData22, "fill:"+lineColor)
	canvas.Path(pathData23, "fill:"+lineColor)
	canvas.Path(pathData24, "fill:"+lineColor)
	canvas.Path(pathData25, "fill:"+lineColor)
	canvas.Path(pathData26, "fill:"+lineColor)
	canvas.Path(pathData27, "fill:"+lineColor)
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
	canvas.Path(pathData60, "fill:"+hexcode)
	canvas.Gend()
	canvas.Gend()
}
