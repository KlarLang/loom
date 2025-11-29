#!/bin/bash 

set -e 

CUSTOM_COLOR='\033[38;2;127;0;31m' 
NORMAL='\033[0m'

show_help() { 
    echo "K - CLI" 
    echo 
    echo "Usage:" 
    echo -e "${CUSTOM_COLOR}k${NORMAL} <command> [options]" 
    echo 
    echo "Commands:" 
    echo -e "${CUSTOM_COLOR}update${NORMAL}    | ${CUSTOM_COLOR}u${NORMAL}    ->   Update the language" 
    echo -e "${CUSTOM_COLOR}lex${NORMAL}       | ${CUSTOM_COLOR}l${NORMAL}    ->   Lexicalizes a file" 
    echo 
    echo "Others:"
    echo -e "${CUSTOM_COLOR}--help${NORMAL}    | ${CUSTOM_COLOR}-h${NORMAL}   ->   Show this help log"
    echo -e "${CUSTOM_COLOR}--version${NORMAL} | ${CUSTOM_COLOR}-V${NORMAL}   ->   Show version" 
    echo 
} 

# No arguments → show help 
if [ $# -eq 0 ]; then 
    show_help 
    exit 0 
fi 

cmd="$1" 
shift 

case "$cmd" in 
    --help | -h)
        show_help 
    ;;

    --version | -V) 
        echo "K CLI -> v0.1.1" 
        echo "Klang -> v0.1.10" 
    ;; 

    update | u) 
        echo "Running update..." 
        # here enter your code later 
    ;; 

    lex | l) 
        file="$1" 
        if [ -z "$file" ]; then 
            echo "Error: You need to provide a file" 
            exit 1 
        fi 
            echo "Running lex on $file" 
            # here you would run "java -jar core.jar lex $file" 
        ;; 

    *) 
        echo "Unknown command: $cmd" 
        echo 
        show_help 
        exit 1 
    ;; 
esac