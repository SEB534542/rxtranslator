# Reactive Resume translator
This module can be used to translate a JSON export of a resume from [rxresu.me](https://rxresu.me) (or self hosted).
The cmd folder contains a program to translate the JSON via the command line with the following arguments:
> rxtranslator <API_KEY> <REACTIVE_RESUME_JSON_FILENAME> <TARGET_LANG>

Where 
* <REACTIVE_RESUME_JSON_FILENAME> refers to the location of the JSON export of the reactive resume, and
* <TARGET_LANG> refers to the target language as specified by [DeepL](https://developers.deepl.com/docs/api-reference/translate/openapi-spec-for-text-translation).
