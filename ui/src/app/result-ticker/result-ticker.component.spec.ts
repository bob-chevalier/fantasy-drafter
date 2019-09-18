import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ResultTickerComponent } from './result-ticker.component';

describe('ResultTickerComponent', () => {
  let component: ResultTickerComponent;
  let fixture: ComponentFixture<ResultTickerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ResultTickerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ResultTickerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
