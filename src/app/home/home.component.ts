import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { NavComponent } from '../nav/nav.component';
import { HttpClient } from '@angular/common/http';
import { HttpClientModule } from '@angular/common/http';
import { FooterComponent } from "../footer/footer.component";

interface Car {
  make: string;
  carmodel: string;
  year: number;
  price: number;
  mileage: number;
  description: string;
  image_url: string;
}

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CommonModule, RouterModule, NavComponent, HttpClientModule, FooterComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
getUniqueYears(): any {
throw new Error('Method not implemented.');
}
  cars: Car[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.fetchCars();
  }

  private fetchCars(): void {
    this.http.get<Car[]>('http://localhost:8080/cars').subscribe({
      next: (cars: Car[]) => {
        this.cars = cars;
      },
      error: (error: any) => {
        console.error('Error fetching cars:', error);
      }
    });
  }
}