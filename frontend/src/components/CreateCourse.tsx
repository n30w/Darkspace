"use client";

import React, { useState } from "react";
import CloseButton from "@/components/buttons/CloseButton";

interface props {
  onClose: () => void;
  onCourseCreate: (courseData: any) => void;
}

const CreateCourse: React.FC<props> = (props: props) => {
  const [courseData, setCourseData] = useState({
    title: "",
    professor: "",
  });

  const postNewCourse = async (courseData: any) => {
    try {
      const res: Response = await fetch("/v1/course/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          title: courseData.title,
          teacherid: courseData.professor
        }),
      });
      if (res.ok) {
      } else {
        console.error("Failed to create course:", res.statusText);
      }
    } catch (error) {
      console.error("Error creating course:", error);
    }
  };

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setCourseData({
      ...courseData,
      [name]: value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    props.onCourseCreate({ ...courseData });
    postNewCourse(courseData);
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            Create New Course
          </h1>
          <div className="mb-2">
            <label
              htmlFor="title"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Course Name:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={courseData.title}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
              required
            />
          </div>
          <button
            type="submit"
            className="w-full inline-flex justify-center mt-8 px-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 py-2 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
          >
            Create
          </button>
        </form>
      </div>
    </div>
  );
};

export default CreateCourse;
