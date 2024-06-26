"use client";

import React, { useState } from "react";
import CloseButton from "@/components/buttons/CloseButton";

interface props {
  onClose: () => void;
  // onAnnouncementCreate: (announcementData: any) => void;
  params: {
    id: string;
  };
  token: string;
}

const CreateAnnouncement: React.FC<props> = (props) => {
  const initialAnnouncement = {
    courseId: props.params.id,
    token: props.token,
    title: "",
    description: "",
  };
  const [announcementData, setAnnouncementData] = useState(initialAnnouncement);

  const handleChange = (e: { target: { name: any; value: any } }) => {
    const { name, value } = e.target;
    setAnnouncementData({
      ...announcementData,
      [name]: value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    const postNewAnnouncement = async (announcementData: any) => {
      const res: Response = await fetch(
        `http://localhost:6789/v1/course/${announcementData.courseId}/announcement/create`,
        {
          method: "POST",
          body: JSON.stringify({
            courseid: announcementData.courseid,
            token: announcementData.token,
            title: announcementData.title,
            description: announcementData.description,
            media: [],
          }),
        }
      );
      return res;
    };

    postNewAnnouncement(announcementData).catch(console.error);
    props.onClose();
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white rounded-lg shadow-lg px-32 py-16 justify-end">
        <CloseButton onClick={props.onClose} />
        <form className="justify-end" onSubmit={handleSubmit}>
          <h1 className="font-bold text-black text-2xl pb-8">
            Create Announcement
          </h1>
          <div className="mb-2">
            <label
              htmlFor="title"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Announcement Title:
            </label>
            <input
              type="text"
              id="title"
              name="title"
              value={announcementData.title}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
            />
          </div>
          <div className="mb-2">
            <label
              htmlFor="description"
              className="block text-lg font-medium text-gray-700 py-2"
            >
              Description:
            </label>
            <input
              type="text"
              id="description"
              name="description"
              value={announcementData.description}
              onChange={handleChange}
              className="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md h-8"
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

export default CreateAnnouncement;
