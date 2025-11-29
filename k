#!/bin/bash 

set -e 

show_help() { 
    echo "K - CLI" 
    echo 
    echo "Usage:" 
    echo " k <command> [options]" 
    echo 
    echo "Commands:" 
    echo "version Show version" 
    echo "update Update the language" 
    echo "run Run a file" 
    echo "lex Lexicalizes a file" 
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
    version) 
        echo "K CLI -> v0.1.0" 
        echo "Klang -> v0.1.10" 
    ;; 

    update) 
        echo "Running update..." 
        # here enter your code later 
    ;; 

    lex) 
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