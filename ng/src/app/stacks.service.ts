import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

export interface StackConfigs {
  Stacks: Array<string>
}

@Injectable({
  providedIn: 'root'
})
export class StacksService {

  private apiUrl = 'http://localhost:8080/api/stacks';

  constructor(private http: HttpClient) {}

  getStacks(): Observable<string[]> {
    return this.http.get<string[]>(this.apiUrl);
  }
}
