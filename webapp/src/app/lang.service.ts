import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map } from 'rxjs/operators';
import Lang from './lang';

@Injectable({
  providedIn: 'root'
})
export class LangService {

  uri = 'http://localhost:3000/api/v1/lang';

  constructor(private http: HttpClient) { }

  createLang(langText: string) {
    const obj = {
      lang: langText
    };
    this.http.post(this.uri, obj)
      .subscribe(res => console.log('Done'));
  }

  getLangs() {
    return this.http.get(this.uri).pipe(
      map((array: any[]) => {
        return array.map(lang => {
          return {
            LangID: lang.ID,
            LangName: lang.lang
          } as Lang;
        });
      }));
  }
}
