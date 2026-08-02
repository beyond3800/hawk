package env

import (
	"bufio"
	"os"
	"strings"
)

type EntryPoint int

const (
	Variable EntryPoint = iota
	Comment
	Empty
)
type Entry struct {
	Type    EntryPoint
	Key     string
	Value   string
	Raw     string
}
type Parser struct {
	Path    string
	Entries []Entry
}

func New(path string) *Parser {
	return &Parser{
		Path: path,
	}
}
func (p *Parser) Load() error {
	file, err := os.Open(p.Path)
	if err != nil{
		return err
	}
	defer file.Close()

	p.Entries = nil

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()
		entry := Entry{
			Raw: line,
		}
		trim := strings.TrimSpace(line)
		switch {
		case trim == "":
			entry.Type = Empty
		case strings.HasPrefix(trim,"#"):
			entry.Type = Comment
		default:
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2{
				entry.Key = strings.TrimSpace(parts[0])
				entry.Value = strings.TrimSpace(parts[1])
			}
		}
		p.Entries = append(p.Entries, entry)
	}
	return scanner.Err()
}

func (p *Parser) Save() error{
	
	var builder strings.Builder
	
	for i, entry := range p.Entries{
		
		switch entry.Type  {
		case  Empty:
			builder.WriteString("")
		case Comment:
			builder.WriteString(entry.Raw)
		default:
			builder.WriteString(entry.Key)
			builder.WriteString("=")
			builder.WriteString(entry.Value)
			
		}
		if i != len(p.Entries)-1{
			// if entry.Type == Variable{
			// 	prefix := string(entry.Key[:3]) 
				
			// 	if strings.HasPrefix(entry.Raw,prefix){
			// 		fmt.Println(prefix)
			// 		builder.WriteString("\n")
			// 	}else{
			// 		builder.WriteString("\n\n")
			// 	}
			// }
			builder.WriteString("\n")
		}
	}
	return os.WriteFile(p.Path,[]byte(builder.String()),0644)
}

func (p *Parser) Get(key string) (string, bool){

	for _, entry := range p.Entries{
		if entry.Key == key{
			return entry.Value, true
		}
	}
	return "", false
}

func (p *Parser) Has(key string) bool{

	_,ok := p.Get(key)

	return ok
}

func (p *Parser) Set(key, value string){
	
	for i, entry := range p.Entries{
		
		if entry.Key == key{
			p.Entries[i].Value = value
			return
		}
	}
	p.Entries = append(p.Entries, Entry{
		Key: key,
		Value: value,
	})
}

func (p *Parser) Delete(key, value string){
	
	var entries []Entry
	for _, entry := range p.Entries{
		
		if entry.Key == key{
			continue
		}
		entries = append(entries, entry)
	}
	p.Entries = entries
}
