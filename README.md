# johnny

![image of ghost rider](https://i.pinimg.com/736x/77/93/c6/7793c64d4fa2d4d32560e978070d90c8.jpg "Johnny Blaze")

## Why?

* You work at Skit.
* You repeatedly download s3 audios, and convert them to 8kHz 16bit .wav audios.
* And you want them fast.


## Usage

You provide a .csv containting a column called `audio_url` / `s3_audio_url`.
And `johnny`, downloads all those audios, converts them to 8kHz 16bit .wav audios, and puts them in a directory you provide in `-output`.


```
$ ./johnny -h
Usage of ./johnny:
  -input string
    	csv which contains audio urls (default "input.csv").
  -output string
    	directory where the wav audios need to be stored. (default "wav_audios")
  -workers int
    	maximum goroutines in the pool (default 30)
```


```
$ ./johnny -input tagged_data.csv -output data/wav_audios/ -workers 50
 100% |████████████████████████████████████████████████████████████| (967/967, 90 it/s)         
finished processing 967 audios and stored them in data/wav_audios/
```

