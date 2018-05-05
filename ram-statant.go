package heraldry

import svg "github.com/ajstarks/svgo"

func renderRamStatantToSvg(canvas *svg.SVG, hexcode string, lineColor string) {
	pathData0 := "M88.9,117.2c15.9-10.9,37.5-3.2,42.4,13.4c0.5,1.8,0.9,3.6,0.6,5.3c-1,5.3,2.6,6.9,6.4,8c17.8,4.8,35.9,3.5,53.9,2.1    c18.5-1.4,35.3,2,49.1,15.6c4.7,4.7,4.8,11,6.3,16.4c2.2,8.4,3.8,17.1,3.6,26.1c-0.1,4.2-1.7,7.6-3.8,10.8c-2,3.1-5.5,4-9.4,2.3    c0,8.3,0.1,16.5,0,24.6c-0.1,8.2-12,15-18.8,10.8c-0.9-0.6-1.8-1.5-0.6-2.9c0.8-1,1.3-2.2,2.2-3.2c7.1-7.7,4.1-16.3,2.3-24.8    c-0.4-2-1-1.1-2.4-1.3c-6.9-0.8-9.3,2.9-12,8.6c-2.2,4.6-4.1,9.4-6.1,14.1c-2.1,4.9-5.8,6.9-10.4,9.1c-4.8,2.2-8.7,0.4-13.1,0.5    c0.5-2.6,2.3-3.6,3.8-4.9c9.2-8,15-17.8,16.3-30.1c0.2-1.4,0.8-2.8,0.9-4.2c0.4-4.9-1.3-6.9-6.1-6.5c-14.4,1.2-28.8,0.8-43.3,1    c-3.2,0.1-6.8-2.4-9-0.6c-2,1.6-0.4,5.7-0.7,8.6c-1.1,8.7-3.1,17.3-3,26.2c0.1,8.9-4.9,14.2-13.4,14.6c-0.3-3.6-5.5,0.1-5.7-3.9    c-0.1-1.6-1.3-1-2.1-0.3c-2.8,2.7-5.5,0.7-7.2-0.9c-2.1-2,0.6-3.4,1.9-4.7c4.1-4.4,4-10,3.4-15.3c-1-8.3-2.5-16.6-4.8-24.7    c-0.3-0.9-0.1-2-1.1-2.6c-8.4-5.2-11.1-13.5-13.3-22.6c-1.3-5.4-5.5-10.1-8.4-15.1c-0.6-1.1-1.5-2.1-2.1-3.3    c-3.4-6.6-11.1-7.9-16.4-2.8c-4.6,4.4-13.3,2.2-15.4-3.7c-0.7-2,0-3.5,0.9-4.4c1.6-1.8,1.4-2.8-0.3-4.2c-1.7-1.5-3.2-2.1-5.3-0.7    c-1.8,1.2-3.1,5.3-5.8,2.4c-2.1-2.3-1.5-5.8-0.1-8.8c2.7-5.3,8.9-7.3,15.5-5.6c1.9,0.5,4.4,2.9,5.5,1.5c2.1-2.8,5.2-5.3,5.1-9.6    c-0.1-5.8,0.1-11.7,2.9-17.1c2.5-4.7,11.4-7.8,15.8-5.2c3.2,1.9,4.2,5,2.8,8.6C90.2,115.1,89.5,116.1,88.9,117.2z"
	pathData1 := "M127,246.9c0.8-10.6,3.1-21.2,0.1-31.9c-0.7-2.3-0.1-5-0.1-7.5c2.8-0.3,4.1,3.4,7.1,2.6    c0.7-0.2,1.1,1.1,0.6,1.4c-2.7,1.2-0.8,2.5,0.2,2.8c2.1,0.7,1.1-1.7,1.5-2.5c1-1.6,1.3-3.9-0.6-5.3c-0.9-0.7-0.3-3.1-2.5-2.2    c-2.8,1.2,0.7,2.4,0.2,3.7c-3.3-0.2-2.2-3.1-3.2-4.7c-0.2-0.3,1.2-1.4,0.1-2c-1.2-0.7-1.7,0.5-2.3,1.2c-0.4,0.5-0.8,1.5-1,1.4    c-1.3-0.3-2.8,1.9-3.8-0.1c-1-2.1,0.6-2.5,2.2-2.8c0.6-0.1,1.4,0,1.4-1c-7.5-0.1-7.7,0.1-5.9,5.9c-1.8-1-3.9-3.2-3.6-3.4    c3.2-2.7-2.7-1.5-1.3-3c1.1-1.4,2.4,1.3,3.4-0.3c0.2-0.4,1.6-0.7,0.9-0.8c-4.4-1.1-0.5-4.1-1.7-6c-1.7,0.4-2.4,1.9-3.1,3.3    c-0.2-2.4-0.1-2.6,2-5.7c-1.6,0.8-2.3,1.3-3.1,1.6c0.6-2.9,0.5-3.1-4-6c-0.8-0.5-0.4-1.4-0.4-2.1c0-1.2-0.3-1.7-1.5-2.8    c-0.8-0.7-0.7-1.4,0.5-1.7c1.7-0.3,0.5-2,1.7-2.5c-2.1-1.5-1.7,2.9-3.4,1.1c-1.1-1.1-0.2-2.4,0.9-3.4c1.1-3.1,1.1-3.2,4.3-2.8    c0.6,0.1,0.6-0.4,0.8-0.7c0.2-0.3,0.2-0.6,0.1-0.9c-0.3-0.8-0.9-0.8-1.6-0.6c-2.7,0.7-3.4,2.8-3.5,5.2c-3.5-0.1-0.4-3-1.7-4.1    c-1,1.3-0.9,1.1-3,0c-1.6-0.8-2.5,0.9-4.6,1.8c7.9,2.5,6.1,9.3,8,14.4c0.1,0.1,0,0.3,0,0.5c-0.4,1.7-4.3,0.6-2.7,3.8    c1.3,2.5,3.5-0.3,5,1.2c1.6,1.7,1.1,4.4,3.3,5.7c-1.8,0.3-1.5,2.4-2.8,3.1c-1.5,0.8-2-0.7-2.8-1.3c-4.8-4.2-7-9.8-8.2-16.2    c2.7,0.7,3.3,3.5,6,3.8c0-2.1-1.6-4.3,0.2-6.2c0.2-0.2,0.4-1.6-0.6-1.4c-4.7,1.1-3.3-2-2.6-4.4c-2.3,0.2-2.8,2.3-4.4,2.8    c-0.8-2.2-3.1-4.2-0.9-6.8c0.7-0.9-0.2-1.2-0.6-1.5c-1.8-1.6-2.3,1.3-4.1,1.2c2.1-3.3,2.3-7.5,6.9-9c1.7-0.6,3.5-1.2,4.7-0.5    c8.2,4.9,13.8-0.8,19.7-4.8c1-0.7,1.3-2.3,2.4-4.3c1.8,5.4-2.9,8.2-0.9,12c5.7-3.8,3.3-10.6,6.2-15.4c3,1.4-2.1,4.6,1.5,5    c1.7,0.2,0.2-2.3,1.2-3.3c1.6,0.7-0.4,3,1.5,3.4c0.9,0.2,0.9-0.6,1.3-1.1c0.8-1.2,1.8-2.2,2.8-3.4c-3.3,0.5-5.1-1.3-7.8-3.9    c4.9,0,8,4,12,2.2c1.1,1.9-2.1,2.3-0.7,4.2c1.5-0.1,2.8-0.9,4-1.9c0.6-0.5,1.2-1.5,2-0.7c0.7,0.6-1.3,1.5-0.2,2.4    c1.8,0,2.8-5.1,5.4-0.8c0.1,0.2,1.8,0.1,2.2-0.4c1.7-1.9,3-2,4,0.5c0.3,0.8,1.2,1.2,1.3,1c1.5-3.8,4.9-1.5,7-1.7    c5.4-0.5,11-0.7,16.3,0.8c3.2,0.9,6-0.9,8.9-1.4c18-2.7,33.6,1.4,46.3,14.8c0.4,0.5,0.8,1,1.1,1.6c0.2,0.4,0,1-0.1,1.8    c-3-3.3-5.4-6.8-9.5-8.6c-0.7,2.1,1.4,2.5,2,3.7c1.6,2.7,2.3,5.9,2.1,8.6c-0.2,4.9-0.8,4.8,1.5,9.2c-2.7-0.4-3.2-4.5-6.6-3.4    c0.9,3.2,7.4,4.5,3,9.6c-0.3-2.7-0.7-5.2-4-6c2.3,3.9,2.3,3.9-2.4,5.1c-0.8,0.2-1.2,1.1-0.8,1.3c2,0.9-0.7,2.8,0.8,3.5    c0.9,0.4,1.8-0.1,2.6-0.9c0.5-0.4,0.2-3.1,2-0.9c0.8,0.9-2.1,2.7,0.6,2.8c1.8,0,1.7-2.1,2.2-3.5c0.7,4,0.7,3.8-3.1,5.1    c-0.9,0.3-4.1-0.1-2.3,2.9c0.4,0.7-0.5,2.2-1.1,1.6c-3.3-3.6-5,2.7-7.9,1.2c-0.1-0.1-0.8,0.5-0.9,0.8c-0.1,0.4,0,0.9,0.7,1    c3.4,0.8,10.3-2.2,12.6-5.6c0,3.4-2,5.8-5,5.8c-0.7,0-1.2,0-1.5,0.8c-0.3,1.1,1.2,3.1,0.6,2.7c-3.9-2.3-7.5,3-11.7,0.1    c1.1-1.4,5.8,0.6,3.2-3.7c-2-3.3-2.4,2.7-4.6,0.3c-2-2.1,3.2-3.6,0-4.1c-1.6-0.3-1.2-3.6-3.2-1.8c-1.4,1.3,0.1,2.5,0.6,3.8    c0.3,0.8,0.5,2-0.6,2.8c-3.7-1.1,0.4-8.8-6.2-7.3c-0.3-2.5,2.9-2.8,2.8-5.5c-3.4,2.9-4.2-0.1-5.7-2.2c-1.2,2-2.3,3.9-3.2,5.5    c-4.3-0.1-4.3-9.4-11.7-3.8c6.2,0.4,7.5,7.9,14.4,6.3c-3.6,2.7-6.5,2-9.4-1.4c-1.2-1.4-5.9-1.8-7.3-0.5c-0.9,0.8-1,0.7-3.3-2.7    c-1.6,0.4-1.9,1.7-1.5,2.9c0.4,1.4,1.7,1.8,3.2,1.7c2.4-0.2,4.9-1.2,5.4,3c0.2,1.5,2.5,3.4,4.7,1.1c0.9-0.9,1.8-1.4,3.3-0.8    c-0.5,2.3-3,1.9-4.3,3.1c0.8,1.5,2.3,1.1,3.6,1.9c-4.1-0.8-7.9,0.1-12,2.4c0.8-2.1,2.6-1.8,2.6-3.2c-0.8-1.3-3.4,0.7-3.4-1.4    c0-1.4,1.1-2.9,3-1.3c0.7,0.6,1.2,1.8,2.4,1.1c-0.9-6.1-6.2-3.1-9.3-4.8c-1.1,2.8,2.8,1.7,2.2,3.8c-1.5,1.2-2.8-1-4.5-0.6    c0.1,4.1,4.7,2,5.4,4.5c-3.7,2.4-5.6-2.3-8.8-2.3c-0.1-0.8,0.3-2.1-1.3-1.3c-0.2,0.1,0.2,0.2,0.4,0.2c0.4,0.2,0.8,0.5,1.1,1    c-3.4,3.3,2.9,1.1,1.5,3.5c-2.6,1.8-4.1-1.4-6.2-2.1c-1.5-0.5-1.1-0.5-2.6,0.8c-1.1,1-2.9,1.8-4.8,1.6c1.4-3.4,2.9-6.6,6.9-8.7    c-2.6-0.4-5.8,2.9-6.2-1.9c0-0.3-3.3,1.3-5.3,0.1c0.5,0.9,0.8,2.1,1.6,2.5c1.6,0.8,5-0.2,3.2,3.6c-2.1,4.3-3.8,0.5-5.7-0.2    c-0.8,1.4-0.2,2.5,0.9,3.5c1.4,1.2-0.2,2-0.7,2.1c-1.9,0.5-0.5-1.1-0.8-1.5c-1.1-1.4-2.5-2.7-3.8-3.9c-0.5-0.4-1.6-0.7-1.6,0    c0.3,5.1-5.6-0.9-5.7,3.2c-0.3-0.2-0.6-0.2-0.8-0.5c-0.3-0.3-0.9-0.9-0.2-1.1c2.6-1,4.3-3.8,7.5-3.5c1.2,0.1,2.9,2,2.3-1.5    c-0.2-1.3,2,0.8,3.6,0.5c-0.9-1.9-1.8-4-2.8-6c-1.9-0.3-0.3,2.5-2,2.1c-1.3-0.3-2-1.3-2.7-2.5c-0.8-1.5-0.1-3-0.6-4.4    c-0.4-1.2-1.5-2.3-2.6-1.6c-2,1.2,1.3,0.7,0.7,1.8c-0.9,1.7-1.8,3.5-0.1,5.3c0.1,0.1-0.7,1.2-1.2,1.4c-0.3,0.1-1.3-0.5-1.4-0.9    c-0.5-1.6-0.9-1.7-2.1-0.5c-1.1,1.1-1,1.6,0.1,2.6c2.4,2.1,4.8,2.1,7.5,0.5c-1.9,3.2-4.5,3.2-7.2,1.6c-1.7,1.7,1.4,6.7-4,5.2    c-0.2-0.1-1.5,2-1.3,2.3c4.4,6.8,0.6,13.9,0.2,20.6c-0.4,6.7,0.8,13.9-2.2,20.7c-0.3-0.9-0.6-1.6-1.1-3c-0.4,2.7-1.6,5.5-1.1,6.4    c2,3.5-1.1,3.4-2.4,3.7c-2.1,0.7-4-0.2-4-2.9c0.7-2.4,5-0.9,4.7-4.1c0.5-0.7-1-0.3-0.3-0.1c0.4,0.1,0.3,0.2,0.2,0.5    C129.7,248,128.5,248.2,127,246.9z"
	pathData2 := "M94.9,125.9c-0.1-3.6-2.5-6.4,0.9-8.8c0.2,2,0.3,3.7,0.5,6.3c1.1-3,0.2-6.1,3.5-7.1c0.4,1.8,0.2,3.8,1.7,5.3    c0.9-0.5,0.6-1.4,0.5-2.1c-0.2-1.6-2.1-3.9,0.8-4.3c2.5-0.4,0.3,2.9,1.9,3.7c2.5-2.3,4.7-7.4,8.5-1.4c3.6-2.9,4.6,1.7,6.9,2.5    c2.8,1,3.8,3.1,1.1,5.6c1.9,0.8,2.1-1.6,3.5-1.3c0.9,1.7-1.5,1.9-1.4,3.3c1.7,1.2,2-1.7,3.6-1.2c0,2.1-2.1,2.5-2.7,4.1    c1.8,1.5,2.1-1.5,4-1.1c1.8,0.3-5.1,3.7,0.6,3.4c0.4,0-2.7,0.8-2.6,2.6c2.9-0.8,1.3,2.4,2.7,3c-3.5,2.7-1.4,6.9-2.9,10.2    c-1.5,3.2-3.6,5.5-7.7,7.5c2.1-2.6,2.6-4.8,4.8-6.4c1.9-1.4,1.4-4.6,1.2-6.6c-0.7-7.1-6.2-17.9-16-18.5c-3.4-0.2-6.1,0.9-8.9,2.1    c-2.2,0.9-4.6,2.9-6.7-0.3c-0.6-0.8-1.6-0.7-1.7,0.4c0,0.6-1.8,1.2,0.1,2.2c2.4,1.3,4.6,2.4,7.8,1.8c-2.1,3.6-5.4,1-7.9,2.7    c3,0.4,5.6,2.2,8.4,0.4c4.8-3.1,9.1-1.9,13,1.6c-1.2,0.5-2.2,2.6-3.6-0.7c-0.9-2-4.5-1-6.8-0.4c-1.4,0.3-2.6,1.3-3.8,2    c0.4,1,1.5,0.3,2.2,0.8c-2.1,3.1-4.1-0.5-6.1-0.1c0.9,10.4-2.5,14.4-12.9,15.1c0,1.5,1.9,0.2,2.1,1.4c-6.2,2.5-14.3-0.8-18.8,6.9    c0-2.6,2.2-3.2,3.1-5.5c-4,0.4-5.6,7.3-11.6,3.5c4-0.4,5.7-1.8,4.4-5.1c-0.2,0.1-0.5,0.1-0.8,0.3c-1.4,0.8-1.3,4.4-3.4,2.9    c-1.6-1.2,0.1-3.4,1.4-5c2.1-2.6,4-5.5,5.5-8.5c2.8-5.5,7.1-10,10.3-15.2c1.6-2.6,5.2-2.8,7.2-0.7c1.1,1.1,1.3,0,1.7-0.4    c1.4-1.2,2.4-2.7,4.3-3.2c0.4-0.1,0.7-0.8,1.1-1c0.8-0.3,2.9-1.5,3.3-2.5C92.3,121.9,92.4,124,94.9,125.9z"
	pathData3 := "M226,207.5c2.4-1.9,4.3-4.2,8.2-4.3c-2.7,2.7-2.6,5.2-1.5,8.2c3,8.4,2.4,17.1,2.3,25.8c0,3.8-4.5,2.7-5.7,5.3    c2.3,2.9,3.4-1.7,5.7-0.9c-0.7,5-4.4,7.3-8.5,9.3c0.8-2,3.1-3.3,3.1-6.2c-3,1.7-3.6,5.9-7.3,5.9c-0.1-0.3-0.3-0.8-0.2-0.9    c6.4-4.7,6.1-11.4,6-18.3c0-7-0.8-13.5-4.7-19.5c-0.5-0.8-1.7-1.8-0.8-2.9c1.3-1.6,1.7,0.4,2.6,0.8c0.3,0.1,0.6,0.2,1,0.2    c1.6,0.4,2.8-0.6,4.2-1.1c1-0.4,0.1-0.8-0.1-1.1C229.1,205.7,227.7,208.6,226,207.5z"
	pathData4 := "M246,188.9c-0.2,1.6-1.8,2.2-2.7,3.4c-0.6,0.9-0.2,1.6,0.5,2.1c1.4,1.1,0.9-0.6,1.3-0.9    c0.7-0.5,2.2-0.1,1.7,0.4c-3.6,3.5,3.4,6.7-0.3,11.4c-2.6,3.3-0.7,4.9-5.2,7.2c1.5,1.8,2.8-0.9,4.5-0.1c-0.9,1.4-2.3,2-3.4,2.9    c-1.8,1.5-3,0.1-3.9-0.5c-1.6-1.1,0.4-1.8,0.8-2.7c1.2-2.4,3.9-4.5,2.1-7.9c-2.4,1.6-2,5.4-5.4,6.5c0.3-5.6,6.4-10.4,2.4-16.4    c-1.1,1.6-0.1,3,0.4,4.6c0.3,1-0.8,2-1.7,2c-1.6,0-0.5-1.2-0.7-1.9c-1.3-7,4-10.8,7.3-15.6c-1.9-1-2.2,1.4-3.6,1.4    c-1.5-3.6,4.7-4.4,3.4-6.7c-0.9,3.2,4.7,8.3-2.6,12c3,0.9,3.1-1.8,4.7-1.7C245.8,188.4,245.9,188.6,246,188.9z"
	pathData5 := "M111.5,159.8c3.4-9.2,2.4-14-3.1-15.7c-3.8-1.2-7.9,0.8-8.5,4.1c-0.8,4.6,2.8,8,9.7,9    c-1.4,3.4-3.7,2.9-6.3,1.5c-6.6-3.3-8.9-7.3-7.3-13c1.4-4.8,5.5-6.2,9.7-6.7c4.8-0.6,10.6,4.3,12.1,9.6    C119,152.5,116.3,157.2,111.5,159.8z"
	pathData6 := "M87,111.8c0,5.8-7.9,13.4-13.4,13.2c-1,0-1.9-0.5-1.4-1.4c2-3.8-0.1-8.5,3.5-12.3c2.9-3.1,5.9-3.6,9.4-3.2    C87.2,108.3,87.1,110.3,87,111.8z"
	pathData7 := "M202.5,222.2c-3.2,3.7,2.3,4.4,2.5,6.9c-0.4,1.1-0.9,0.5-1.3,0c-0.4-0.3-0.7-0.7-1.1-1    c-0.5,1.3-2.2-0.5-2.5,0.8c-0.3,1.1-0.3,2.5,1.6,1.3c0.1-0.1,1.2,1.2,1.1,1.4c-1.9,3.1-0.9,7.5-4,9.9c-1.2,0.9-2.7,1.9-4.5,1.5    c-1.6-0.4-2.9-0.7-1.1-2.8c1.7-2.1,2.9-4.6,4.4-7c0.6,1.5,1.2,2.9,1.8,4.3c-1.7-8,0.9-15.5,3-23.6c1.4,3.2,2.7,6,3.9,8.9    c0.7-0.9,1.3-1.7,2.4-3.1c0.4,3-1.1,5.5-2.4,6.1C204.2,226.7,204.4,223.3,202.5,222.2z"
	pathData8 := "M117.2,223.1c3,2,4,5,4.9,8.1c-2.1,0.9-0.6-2.7-3.2-1.1c-0.2,0.3-0.1,1.2,0.7,2.1c3.7,4.4,2.5,9.6,2.3,14.5    c-0.1,3-3.7,2.5-5.4,4.2c-1.7-2.1,2.6-2.6,1.2-4.6c-2.4,1-3.4,4.2-6.8,4.5C122.2,243.5,117.7,233,117.2,223.1z"
	pathData9 := "M44.3,147.9c-1.2-4.7,1.9-7.3,4.8-9.3c4.2-2.9,8.5-1.1,12.5,1.4c-2.1,3-2.1,3-3.9,2.3    c-7.8-2.8-9.3-2.1-12.5,4.9C45.1,147.4,44.7,147.6,44.3,147.9z"
	pathData10 := "M91.6,168c-1.6-2.2-1.3-2.2-0.8-4.8c0.7-3.3-2.4-3.8-4.6-5.3c3.3,0.1,6.5,1.6,9.8,0.2c0.9-0.4,1,1.1,0.9,1.8    c0,0.7,0.3,2.1-0.9,1.9C90.2,160.5,92.8,165.4,91.6,168z"
	pathData11 := "M214.7,213.7c0.2-0.7,0.5-1.6,1-3c0.7,2.8-0.1,5,1.1,7.1c-5.5-1-8.3-6-6.2-10.6c1.6,0.5,1.7,2.2,2.4,3.4    C213.4,211.7,212.9,213.4,214.7,213.7z"
	pathData12 := "M114.5,133.4c3.8-0.3,8.6,6.4,7.4,10.7c-0.3,1.2-1.1,2.3-1.1,4.6C118.3,143,119.8,136.8,114.5,133.4    L114.5,133.4z"
	pathData13 := "M186.5,251.9c3.3-6.6,3.3-8.1,10-7C193.9,248.4,189.6,249.2,186.5,251.9z"
	pathData14 := "M101,149.3c5.3-1,6,4.3,9.7,4c0.5,0,0.3,1.2-0.3,1.5c-1.4,0.8-3.2,1-4.2-0.3c-0.6-0.8-4.1,0-1.8-2.7    C103.6,151.2,102.7,150.6,101,149.3z"
	pathData15 := "M122.9,214.6c-3.9-2.6-5.1-7.3-8.5-10.3C119.5,205.8,120.7,210.6,122.9,214.6z"
	pathData16 := "M127,246.9c4.6,1.2,2.2-3.2,3.5-4.7c1.4,4.5,0.9,7.6-4.6,7.8c0,0,0.1,0.1,0.1,0.1c-0.7-0.6-0.6-1.3-0.2-2    C126.3,247.7,126.7,247.3,127,246.9z"
	pathData17 := "M117.7,213.3c3.8,1.6,5.1,4.5,5,8.4c-1.4-3-4.2-5.1-4.9-8.5L117.7,213.3z"
	pathData18 := "M204.1,213.1c0.7,0,1.3-0.1,1.3,0c0.8,2,3.1,3.2,3.1,5.7c0,0.1-1.2,0.3-1.3,0.2    C206,217.1,204.5,215.5,204.1,213.1z"
	pathData19 := "M120.2,221.2c-2.6-1.5-4.4-3.8-4.5-7.1C117.2,216.5,118.7,218.8,120.2,221.2z"
	pathData20 := "M114.6,133.4c-0.4-1.5-2.3-1.1-2.9-2.2c-0.2-0.4-0.6-0.6-0.1-1.2c0.7-0.9,1.8-1.4,2.5-0.9    c1.5,1-0.7,0.8-0.5,1.2c0.5,1.5,2.2,0.2,3,1C115.9,132,115.2,132.7,114.6,133.4C114.5,133.4,114.6,133.4,114.6,133.4z"
	pathData21 := "M117.9,213.1c-1.5-1.7-4.3-2.6-3.3-5.8c0.8,2.1,3.9,2.9,3.2,6C117.7,213.3,117.9,213.1,117.9,213.1z"
	pathData22 := "M218.1,219c0-2.9,0-4.3,0-5.8C221.5,214.5,219.2,216.1,218.1,219z"
	pathData23 := "M125.9,248c0.1,0.7,0.1,1.3,0.2,2c-1-0.2-1.4,1.5-2.6,0.8C123.5,249.2,124.9,248.8,125.9,248z"
	pathData24 := "M108.4,174.4c-1.7-4.6,2.2-5.8,4.8-7.8c0.9-0.7,1,0.6,0.8,1c-0.7,1.8,3.5,1.4,1.6,3.1c-1.4,1.2-2.3,4.1-5.2,3.2    c-0.3-1.4,2.2-1.7,1.4-3.4c-2.3,0.2-2.4,2.5-3.6,3.8C108.2,174.2,108.4,174.4,108.4,174.4z"
	pathData25 := "M218.1,153.7c4.2,1,5.9,4.3,7.8,7.1c0.3,0.5,0.3,2.1-0.6,2c-1.5-0.3-2.2-2.3-4-2.2C222.1,157.6,218.9,156.6,218.1,153.7z"
	pathData26 := "M119,187.9c-2.5-4.9-1.8-7.9,2.5-9.8C118.7,180.6,121.2,184.3,119,187.9z"
	pathData27 := "M110.4,197.1c-0.7,0.6,0.1,1.7-1.2,1.5c-2.3-0.4-2.4-2.8-3.9-3.9c0.5-0.5,1-1.1,1.5-1.6    C108,194.4,109.2,195.8,110.4,197.1z"
	pathData28 := "M172.8,201.1c-2.2-0.2-4.5-0.1-6.2-2.6c3-0.5,5.7,1.9,8.2-0.4c0.6,2.4-2.1,1.7-2.1,3.2    C172.7,201.3,172.8,201.1,172.8,201.1z"
	pathData29 := "M176.3,190.8c-5.4,1.4-5.8,0.9-4.4-5.5C172.9,187.9,171.6,191.9,176.3,190.8z"
	pathData30 := "M218.3,175.2c1.7-2,0.8-4.2,2.9-5c0.3-0.1,1.5,0.9,1.4,1.1C222.1,172.8,221.9,174.7,218.3,175.2z"
	pathData31 := "M187,170.8c-0.2,0.8,0.4,2.2-1.2,2.1c-0.8-0.1-1.6-0.8-1.4-1.6c0.2-0.8,0.9-2,1.6-2.2C187,168.8,187,170,187,170.8z"
	pathData32 := "M176.7,195.4c-0.8,2-2.5,1.5-4,1.6c-0.8,0-1.8,0.1-1.6-1.1c0.2-1.2,1.3-1.1,1.8-0.4C174.2,197.2,175.4,194.1,176.7,195.4z    "
	pathData33 := "M155,181c-0.2,1-0.9,1.7-1.9,1.9c-0.8,0.2-1-0.5-1.1-1.2c-0.1-1.2,0.3-1.9,1.6-1.7C154.2,180.1,154.9,180,155,181z"
	pathData34 := "M224.4,179.6c0.8,0.2-0.1-3.3,2.2-2.3c0.3,0.1,0.4,0.9,0.3,1.4C226.5,181.3,224.7,179.6,224.4,179.6z"
	pathData35 := "M150.8,182.9c-2-1-2.1-2.4-2.3-3.7c0.5,0,1-0.3,1.4-0.1C152,179.8,150.6,181.4,150.8,182.9z"
	pathData36 := "M208.7,154.6c-1.2,0.6-2.4,0.7-3.6-0.2c-0.1-0.1,0.1-1.1,0.3-1.2C206.9,152.7,208.1,153,208.7,154.6z"
	pathData37 := "M132.7,228.9c0.8-1.1-0.7-3.1,1.3-3.8c0.2-0.1,0.8,0.2,0.8,0.4C135,227,134,228,132.7,228.9z"
	pathData38 := "M220.3,164.7c1.6-0.8,3.1-1,4.6-0.2c0.1,0.1-0.1,1.2-0.3,1.2C222.9,166.4,222.1,164,220.3,164.7z"
	pathData39 := "M120.1,167.1c0-0.2,0-0.4,0.1-0.5c1.1-0.6,0.3-2.8,2.2-2.5c0.2,0,0.6,0.5,0.5,0.7c-0.2,1.8-1.3,2.4-3,2.2    C119.9,166.9,120.1,167.1,120.1,167.1z"
	pathData40 := "M133.2,220.1c1,0.8,2.4,1.2,1.4,2.7c-0.3,0.4-1.1,1.3-1.3,0.7C132.9,222.5,133.2,221.3,133.2,220.1z"
	pathData41 := "M119.9,166.9c0.2,1.8-0.4,2.8-2.7,3.1c0.8-1.4,0.4-3.4,2.9-2.9C120.1,167.1,119.9,166.9,119.9,166.9z"
	pathData42 := "M226.7,173c1.1,0,1.6,0.6,1.9,1.4c0,0-0.6,0.5-0.9,0.6c-0.7,0.1-1.7,0.2-1.7-1C226.1,173.6,226.5,173.3,226.7,173z"
	pathData43 := "M224,168c-0.3,1.3-1.2,1-1.9,1c-0.2,0-1.2,0.1-0.6-0.7c0.4-0.7,1-1.3,1.9-1.1C223.6,167.2,223.8,167.7,224,168z"
	pathData44 := "M87.7,137.6c-1.3,0.9-3.1-0.1-4.4,2.2c-1.7,3.2-6.6-0.3-9.1,2.8c-0.5-0.5-1.3-0.7-0.6-1.8    C77.8,134.4,80.5,133.7,87.7,137.6z"
	pathData45 := "M83.1,129.2c1.4-4.1,4.1-3.3,6.7-2.8C88.3,129.1,85.5,128,83.1,129.2z"
	pathData46 := "M74.3,131.6c0.1-0.2,0.2-0.6,0.4-0.7c1.6-0.9,3.3-1.7,4.9-2.6c0.1,0.3,0.4,0.6,0.3,0.8C79,131.9,76.1,130.5,74.3,131.6z"
	pathData47 := "M231.6,217.2c0.9,1.7,0.2,2.9-1.3,3.7c-0.3,0.2-1.2-0.2-1.2-0.4C229,218.7,230.5,218.1,231.6,217.2z"
	pathData48 := "M232.8,221c0.4,2-0.9,2.7-1.3,3.9C230.4,223.1,231.2,222,232.8,221z"
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
	canvas.Path(pathData16, "fill:"+lineColor)
	canvas.Path(pathData17, "fill:"+hexcode)
	canvas.Path(pathData18, "fill:"+hexcode)
	canvas.Path(pathData19, "fill:"+hexcode)
	canvas.Path(pathData20, "fill:"+hexcode)
	canvas.Path(pathData21, "fill:"+hexcode)
	canvas.Path(pathData22, "fill:"+hexcode)
	canvas.Path(pathData23, "fill:"+hexcode)
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

}