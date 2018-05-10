The .gif format only allows specifing frame time down to the 1/100th of a second.
But what if you want a gif with 100 frames to take 2.5 seconds?  You can't just set each frame to 0.025 seconds.
A good solution would be to use a video format.
A worse solution would be to use this program to time the frames, which for this example would alternate between 0.02 and 0.03 seconds per frame.
  
To make a gif take some time that's not a multiple of 0.01 seconds, make n repititions of the gif take n * desired time.
For example, to make a gif take 0.375 seconds, use `-r 2 -t 75`.  This makes 2 repititions take 0.75 seconds (2 * 0.375).
Really, the first will take 0.37 and the second will take 0.38, but the average time will be what's desired.

### flags
#### -i
input file  
#### -o
output file  
#### -r
repititions of the gif (default 1)  
Copy the original gif this many times in the output file.
If the original file had 5 frames ABCDE, and you pass -r 3, the output file will have 15 frames ABCDEABCDEABCDE.  
This is not the number of times the gif will loop, that value is unchanged from the original.
#### -t
total time in 1/100 seconds (default 100 (1 second))  
The total time that -r repititions of the original gif file take.
If the resulting gif would be above 50 fps, frames are dropped so that each frame takes at least 1/50th of a second.
#### -f
overwrite output file if it already exists  
