import { getInitial, getColor, formatDate } from "./utils";

const Message = ({ speaker, text, timestampMs }) => {
  const initial = getInitial(speaker);
  const color = getColor(speaker);
  return (
    <div className="py-2 flex">
      <div className="avatar h-100 pl-2">
        <div
          className={`w-8 h-8 ${color} rounded-full flex justify-center items-center mr-4`}
        >
          <span className="text-white font-semibold">{initial}</span>
        </div>
      </div>
      <div className="details h-max">
        <div className="timestamp-name text-black flex gap-3">
          <div className=" text-xs text-gray-500">{speaker}</div>
          <div className="text-xs text-gray-600">
            {" "}
            {formatDate(timestampMs)}
          </div>
        </div>
        <div className="texts text-gray-800 text-sm ">{text}</div>
      </div>
    </div>
  );
};

export default Message;
