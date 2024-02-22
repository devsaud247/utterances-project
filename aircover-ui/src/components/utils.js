// Convert timestamp to readable format
export const formatDate = (timestampMs) => {
  const date = new Date(timestampMs);
  return date.toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    hour12: true,
  });
};

// Function to get the first letter of the speaker's name
export const getInitial = (name) => {
  return name.charAt(0).toUpperCase();
};

// Function to generate a consistent color based on the speaker's name
export const getColor = (name) => {
  const colors = [
    "bg-red-500",
    "bg-blue-500",
    "bg-green-500",
    "bg-yellow-500",
    "bg-purple-500",
  ];
  let sum = 0;
  for (let i = 0; i < name.length; i++) {
    sum += name.charCodeAt(i);
  }
  return colors[sum % colors.length];
};
