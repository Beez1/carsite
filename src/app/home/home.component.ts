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