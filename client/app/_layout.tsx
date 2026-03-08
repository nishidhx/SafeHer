import { Stack } from "expo-router";
import { StatusBar } from "react-native";

export default function RootLayout() {
  return <Stack >
    <Stack.Screen name="index" options={{ headerShown: false }} />
    <StatusBar hidden={false} />
  </Stack>;
}
