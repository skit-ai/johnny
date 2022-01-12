# johnny

![image of ghost rider](https://i.pinimg.com/736x/77/93/c6/7793c64d4fa2d4d32560e978070d90c8.jpg "Johnny Blaze")

## Why?

* You work at Skit.
* You repeatedly download s3 audios, and convert them to 8kHz 16bit .wav audios.
* And you want them fast.


## Setup / Installation

1. Head to [releases](https://github.com/skit-ai/johnny/releases) section.
2. Download the binary appropriate for your OS: like Linux & x86, it should either be a `.tar.gz` or `.zip` file.
3. Untar/unzip it, and get started with using it!

## Usage

You provide a .csv containting a column called `audio_url` / `s3_audio_url`.
And `johnny`, downloads all those audios, converts them to 8kHz 16bit .wav audios, and puts them in a directory/path you provide in `-output`.

The audios on S3 can be .raw/.flac or any format. doesn't matter.

Optionally you can provide the resulting output frequency/sample rate you want at `-rate`. ("8k", "16k", "22k", "44k" anything that `ffmpeg` accepts)
Optionally you can also provide the number of concurrent goroutine workers you want at `-workers`.


```
$ ./johnny -h
Usage of ./johnny:
  -input string
    	csv which contains audio urls (default "input.csv")
  -output string
    	directory where the wav audios need to be stored. (default "wav_audios")
  -rate string
    	audio sample rate / frequency of output audios. (default "8k")
  -workers int
    	maximum goroutines in the pool (default 30)
```


```
$ ./johnny -input tagged_data.csv -output wav_audios/ -workers 30
⠏ downloading & converting audios ... (8135/-, 81 it/s) 
->> johnny took 1m40.567284374s for 8135 audios. they are stored under the directory: wav_audios.
```

## Depends on

* [ffmpeg](https://www.ffmpeg.org/)


