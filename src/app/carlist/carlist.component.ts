import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { NavComponent } from '../nav/nav.component';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { FooterComponent } from '../footer/footer.component';

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
  selector: 'app-carlist',
  standalone: true,
  imports: [CommonModule, RouterModule, NavComponent, HttpClientModule, FooterComponent],
  templateUrl: './carlist.component.html',
  styleUrl: './carlist.component.css'
})
export class CarlistComponent implements OnInit {
  cars: Car[] = [];

  constructor(private http: HttpClient) {}
  ngOnInit(): void {
      this.http.get<Car[]>('http://localhost:8080/cars').subscribe({
        next: (data) => {
          this.cars = data;
          console.log('Cars fetched:', this.cars);
        },
        error: (error) => {
          console.error('Error fetching cars:', error);
        }
      });
  } 
}