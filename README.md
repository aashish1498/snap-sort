# Snap Sort

# Usage

To run, use:
```sh
go run .
```

You will be prompted for an input directory (where your media currently is), and an output (where you want them to be stored).

Files will be **copied** to the output folder to ensure originals are retained.

# Output

The files will be stored in folders based on their datetime in the format

```
<output_directory>/yyyy/mm/dd/<file_name>
```

For example:

```
<output_directory>
├── 2023
│   ├── 01
│   │   ├── 12
│   │   │   ├── img_1
│   │   │   └── img_2
│   │   ├── 13
│   │   │   └── img_3
├── 2024
│   ├── 02
│   │   └── 08
│   │       └── img_4
│   ├── 03
│   │   └── 07
│   │       └── img_5
```

# Logic

## Date
SnapSort will attempt to find the date of your picture in the following order:
1. Checking if the picture has a valid `DateTime` EXIF tag
2. Using the file's last modified time
