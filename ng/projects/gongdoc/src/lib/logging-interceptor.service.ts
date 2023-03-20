import { Injectable } from '@angular/core';
import {
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
  HttpResponse,
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable()
export class LoggingInterceptorService implements HttpInterceptor {
  intercept(
    request: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    console.log(`Sending ${request.method} request to ${request.url}`);
    console.log(`Request headers:`, request.headers);
    console.log(`Request body:`, request.body);

    return next.handle(request).pipe(
      tap((event) => {
        if (event instanceof HttpResponse) {
          console.log(`Received response with status ${event.status}`);
          console.log(`Response headers:`, event.headers);
          console.log(`Response body:`, event.body);
        }
      })
    );
  }
}
