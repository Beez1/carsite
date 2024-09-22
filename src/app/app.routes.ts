import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { CarlistComponent } from './carlist/carlist.component';
import { CardetailComponent } from './cardetail/cardetail.component';
import { ContactComponent } from './contact/contact.component';
import { NavComponent } from './nav/nav.component';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'cars', component: CarlistComponent },
  { path: 'car/:id', component: CardetailComponent }, 
  { path: 'contact', component: ContactComponent },
  { path: 'nav', component: NavComponent },
  { path: '**', redirectTo: '' }
];
