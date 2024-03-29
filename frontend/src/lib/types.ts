export interface Announcement {
  id: string;
  title: string;
  date: string;
  description: string;
}

export interface Assignment {
  id: string;
  title: string;
  duedate: string;
  description: string;
}

export interface Course {
  id: string;
  title: string;
  professor: string;
  location: string;
}

export interface Discussion {
  name: string;
}
