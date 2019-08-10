import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class TranslationService {

  uri = 'http://localhost:3000/api/v1/translate';

  constructor(private httpClient: HttpClient) { }

  getTranslations() {
    return this.httpClient.get(this.uri);
  }

  addTranslation(wordID1, wordID2) {
    const obj = {
      wordID1: Number(wordID1),
      wordID2: Number(wordID2)
    };

    return this.httpClient.post(this.uri, obj);
  }

  deleteTranslation(translationID: number) {
    return this.httpClient.delete(`${this.uri}/${translationID}`);
  }
}
