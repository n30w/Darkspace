import Image from "next/image";

const DashboardDiscussions: React.FC = () => {
  return (
    <div className="w-60 h-70 px-4">
      <h1 className="text-white font-bold text-xl border-b-2 border-white mb-4 pb-4">
        Discussions
      </h1>
      <h2 className="text-white text-m font-bold mb-1">
        Balancing Innovation and Maintenance
      </h2>
      <div>
        <Image src={""} alt={""} />
        <h3 className="text-white text-sm mb-2">
          Neo Alabastro posted on Feb 9, 2024 9:16 PM
        </h3>
      </div>
      <p className="text-white text-sm font-light">In my experience, ...</p>
    </div>
  );
};

export default DashboardDiscussions;
