# QR Code Generator and Reader

This is a simple Go application that allows you to create and read QR codes in your terminal. The application provides options to generate QR codes in different formats (PNG, JPG, SVG) and to read QR codes from files. You can also visualize the QR code directly in the terminal in ASCII format.

## Features

- Generate QR codes in **PNG**, **JPG**, or **SVG** formats.
- Save QR codes to a specified folder.
- Read and decode QR codes from images.
- Display the QR code content and ASCII representation in the terminal.

## Requirements

- Go 1.16 or higher
- [skip2/go-qrcode](https://github.com/skip2/go-qrcode) for generating and displaying QR codes
- [tuotoo/qrcode](https://github.com/tuotoo/qrcode) for reading QR codes from files

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/pratamarizki22/qrcode-generator-reader.git
    ```

2. Navigate to the project directory:

    ```bash
    cd qrcode-generator-reader
    ```

3. Install dependencies:

    ```bash
    make deps
    ```

## Usage

### Generate a QR Code

1. Run the application:

    ```bash
    make run
    ```

2. Choose the option to create a QR code.
3. Input the text that you want to encode in the QR code.
4. Specify the file name and format for saving the QR code (PNG, JPG, or SVG).
5. The QR code will be saved in the specified folder and displayed in the terminal in ASCII format.

### Read a QR Code

1. Run the application:

    ```bash
    make run
    ```

2. Choose the option to read a QR code.
3. Input the file path of the QR code image.
4. The application will read the QR code and display the content and its ASCII representation in the terminal.

## Build

To build the project and create the binary:

```bash
make build
```

## Clean

To clean up generated binaries and QR code files:

```bash
make clean
```

## Dependencies

- [skip2/go-qrcode](https://github.com/skip2/go-qrcode) for QR code generation
- [tuotoo/qrcode](https://github.com/tuotoo/qrcode) for QR code reading

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
