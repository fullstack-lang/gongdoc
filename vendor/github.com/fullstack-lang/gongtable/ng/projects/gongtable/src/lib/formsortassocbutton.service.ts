// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { FormSortAssocButtonDB } from './formsortassocbutton-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class FormSortAssocButtonService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  FormSortAssocButtonServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private formsortassocbuttonsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.formsortassocbuttonsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/formsortassocbuttons';
  }

  /** GET formsortassocbuttons from the server */
  getFormSortAssocButtons(GONG__StackPath: string): Observable<FormSortAssocButtonDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<FormSortAssocButtonDB[]>(this.formsortassocbuttonsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched formsortassocbuttons')),
        catchError(this.handleError<FormSortAssocButtonDB[]>('getFormSortAssocButtons', []))
      );
  }

  /** GET formsortassocbutton by id. Will 404 if id not found */
  getFormSortAssocButton(id: number, GONG__StackPath: string): Observable<FormSortAssocButtonDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.formsortassocbuttonsUrl}/${id}`;
    return this.http.get<FormSortAssocButtonDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched formsortassocbutton id=${id}`)),
      catchError(this.handleError<FormSortAssocButtonDB>(`getFormSortAssocButton id=${id}`))
    );
  }

  /** POST: add a new formsortassocbutton to the server */
  postFormSortAssocButton(formsortassocbuttondb: FormSortAssocButtonDB, GONG__StackPath: string): Observable<FormSortAssocButtonDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<FormSortAssocButtonDB>(this.formsortassocbuttonsUrl, formsortassocbuttondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted formsortassocbuttondb id=${formsortassocbuttondb.ID}`)
      }),
      catchError(this.handleError<FormSortAssocButtonDB>('postFormSortAssocButton'))
    );
  }

  /** DELETE: delete the formsortassocbuttondb from the server */
  deleteFormSortAssocButton(formsortassocbuttondb: FormSortAssocButtonDB | number, GONG__StackPath: string): Observable<FormSortAssocButtonDB> {
    const id = typeof formsortassocbuttondb === 'number' ? formsortassocbuttondb : formsortassocbuttondb.ID;
    const url = `${this.formsortassocbuttonsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<FormSortAssocButtonDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted formsortassocbuttondb id=${id}`)),
      catchError(this.handleError<FormSortAssocButtonDB>('deleteFormSortAssocButton'))
    );
  }

  /** PUT: update the formsortassocbuttondb on the server */
  updateFormSortAssocButton(formsortassocbuttondb: FormSortAssocButtonDB, GONG__StackPath: string): Observable<FormSortAssocButtonDB> {
    const id = typeof formsortassocbuttondb === 'number' ? formsortassocbuttondb : formsortassocbuttondb.ID;
    const url = `${this.formsortassocbuttonsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<FormSortAssocButtonDB>(url, formsortassocbuttondb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated formsortassocbuttondb id=${formsortassocbuttondb.ID}`)
      }),
      catchError(this.handleError<FormSortAssocButtonDB>('updateFormSortAssocButton'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in FormSortAssocButtonService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("FormSortAssocButtonService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
      console.log(message)
  }
}