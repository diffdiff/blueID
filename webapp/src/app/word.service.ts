import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import Word from './word';

@Injectable({
  providedIn: 'root'
})
export class WordService {

  uri = 'http://localhost:3000/api/v1/word';

  constructor(private http: HttpClient) { }

  getAllWords() {
    return this.http.get(this.uri).pipe(map((array: any[]) => {
      return array.map(word => {
        return {
          WordID: word.ID,
          LangID: word.LangID,
          WordText: word.word,
          Example: word.example,
          Description: word.description
        } as Word;
      });
    }));
  }

  addWord(wordText, langID, example, description) {
    const obj = {
      word: wordText,
      LangID: Number(langID),
      example,
      description
    };

    return this.http.post(this.uri, obj);
  }
}

