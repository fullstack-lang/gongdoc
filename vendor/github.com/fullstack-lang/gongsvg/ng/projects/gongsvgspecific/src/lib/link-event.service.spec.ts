import { TestBed } from '@angular/core/testing';

import { LinkEventService } from './link-event.service';

describe('LinkEventService', () => {
  let service: LinkEventService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LinkEventService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
