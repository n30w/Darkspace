"use client";

import Image from "next/image";
import Courses from "./components/Courses";
import React, { useState } from "react";
import CreateCourse from "./components/CreateCourse";

export default function Home() {
  const currentDate = new Date().toLocaleDateString();
  const [isCreatingCourse, setIsCreatingCourse] = useState(false);
  const [courses, setCourses] = useState<Course[]>([]);

  interface Course {
    title: string;
    professor: string;
    location: string;
  }

  const handleCreateCourse = (courseData: any) => {
    setCourses([...courses, courseData]);
  };

  const handleClick = (e: any) => {
    setIsCreatingCourse(true);
  };

  const courseDisplay = courses.map((course) => {
    return (
      <Courses
        coursename={course.title}
        professor={course.professor}
        loc={course.location}
      />
    );
  });

  return (
    <div>
      <nav
        style={{
          backgroundImage: `url('/backgrounds/dashboard-bg.jpeg')`,
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      >
        <div className="relative">
          <div className="absolute inset-0 bg-black opacity-70"></div>
          <div className="py-8 px-32">
            <div className="flex items-center gap-4">
              <Image
                src="/backgrounds/NYU-logo.png"
                width="150"
                height="39"
                alt="NYU Logo"
                className="z-10"
              />
              <Image
                src="/backgrounds/darkspace.png"
                width="200"
                height="39"
                alt="Darkspace Logo"
                className="z-10"
              />
            </div>
          </div>
        </div>
      </nav>
      <div className="bg-white bg-cover bg-no-repeat">
        <div className="flex items-center justify-between py-8 px-32">
          <h1 className="font-bold text-2xl">Spring 2024</h1>
          <button
            className="rounded-full bg-black text-white font-light px-4 py-2 h-12"
            onClick={handleClick}
          >
            + Create Course
          </button>
        </div>
        {courseDisplay}
        {isCreatingCourse && (
          <CreateCourse
            onClose={() => {
              setIsCreatingCourse(false);
            }}
            onCourseCreate={handleCreateCourse}
          />
        )}
      </div>
    </div>
  );
}
