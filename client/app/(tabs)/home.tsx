import React, { useState } from "react";
import {
  StyleSheet,
  Text,
  View,
  ScrollView,
  TouchableOpacity,
  Dimensions,
} from "react-native";
import { Ionicons } from "@expo/vector-icons";
import SOSButton from "@/components/SOSButton";

const { width } = Dimensions.get("window");

const FEATURES = [
  {
    icon: "map-outline",
    title: "Safe Route",
    desc: "AI-powered risk scoring for safest path",
    color: "#4ade80",
    badge: "AI",
  },
  {
    icon: "flame-outline",
    title: "Heatmap",
    desc: "Real-time unsafe area visualization",
    color: "#f97316",
    badge: "LIVE",
  },
  {
    icon: "chatbubble-ellipses-outline",
    title: "Safety Assistant",
    desc: "AI guide for stressful situations",
    color: "#a78bfa",
    badge: "AI",
  },
  {
    icon: "warning-outline",
    title: "Report Incident",
    desc: "Community safety reporting",
    color: "#fb923c",
    badge: null,
  },
  {
    icon: "moon-outline",
    title: "Night Mode",
    desc: "Auto-protection after dark",
    color: "#60a5fa",
    badge: "AUTO",
  },
  {
    icon: "location-outline",
    title: "Live Location",
    desc: "Share location with contacts",
    color: "#34d399",
    badge: null,
  },
];

const TIPS = [
  "Stay in well-lit areas at night",
  "Share your live location with trusted contacts",
  "Trust your instincts — if something feels off, leave",
  "Keep emergency contacts updated",
];

export default function Home() {
  const [nightMode, setNightMode] = useState(false);
  const tipIndex = new Date().getHours() % TIPS.length;

  return (
    <View style={styles.container}>
      <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={styles.scroll}>

        {/* Header */}
        <View style={styles.header}>
          <View>
            <Text style={styles.greeting}>Good {new Date().getHours() < 12 ? "Morning" : new Date().getHours() < 17 ? "Afternoon" : "Evening"} 👋</Text>
            <Text style={styles.appName}>SafeHer</Text>
            <Text style={styles.tagline}>Your AI-powered safety companion</Text>
          </View>
          <TouchableOpacity
            style={[styles.nightBtn, nightMode && styles.nightBtnActive]}
            onPress={() => setNightMode(!nightMode)}
          >
            <Ionicons name="moon" size={20} color={nightMode ? "#000" : "#60a5fa"} />
          </TouchableOpacity>
        </View>

        {/* Risk Status Card */}
        <View style={styles.riskCard}>
          <View style={styles.riskLeft}>
            <Text style={styles.riskLabel}>AREA RISK LEVEL</Text>
            <Text style={styles.riskValue}>Moderate</Text>
            <Text style={styles.riskDesc}>2 incidents reported nearby</Text>
          </View>
          <View style={styles.riskScore}>
            <Text style={styles.riskNumber}>42</Text>
            <Text style={styles.riskMax}>/100</Text>
          </View>
        </View>

        {/* Safety Tip */}
        <View style={styles.tipCard}>
          <Ionicons name="bulb-outline" size={18} color="#fad047" />
          <Text style={styles.tipText}>{TIPS[tipIndex]}</Text>
        </View>

        {/* Features Grid */}
        <Text style={styles.sectionTitle}>Features</Text>
        <View style={styles.grid}>
          {FEATURES.map((f, i) => (
            <TouchableOpacity key={i} style={styles.featureCard} activeOpacity={0.8}>
              <View style={styles.featureTop}>
                <View style={[styles.iconBox, { backgroundColor: f.color + "22" }]}>
                  <Ionicons name={f.icon as any} size={22} color={f.color} />
                </View>
                {f.badge && (
                  <View style={[styles.badge, { backgroundColor: f.color + "33" }]}>
                    <Text style={[styles.badgeText, { color: f.color }]}>{f.badge}</Text>
                  </View>
                )}
              </View>
              <Text style={styles.featureTitle}>{f.title}</Text>
              <Text style={styles.featureDesc}>{f.desc}</Text>
            </TouchableOpacity>
          ))}
        </View>

        {/* Night Mode Banner */}
        {nightMode && (
          <View style={styles.nightBanner}>
            <Ionicons name="moon" size={20} color="#60a5fa" />
            <View style={{ marginLeft: 12, flex: 1 }}>
              <Text style={styles.nightBannerTitle}>Night Mode Active</Text>
              <Text style={styles.nightBannerDesc}>Location sharing & heightened alerts enabled</Text>
            </View>
            <View style={styles.nightDot} />
          </View>
        )}

        {/* Recent Activity */}
        <Text style={styles.sectionTitle}>Recent Activity</Text>
        <View style={styles.activityCard}>
          {[
            { icon: "alert-circle-outline", color: "#f97316", text: "Unsafe area marked 0.3km away", time: "2m ago" },
            { icon: "checkmark-circle-outline", color: "#4ade80", text: "Safe route calculated", time: "1h ago" },
            { icon: "people-outline", color: "#a78bfa", text: "Community report verified", time: "3h ago" },
          ].map((item, i) => (
            <View key={i}>
              <View style={styles.activityRow}>
                <Ionicons name={item.icon as any} size={20} color={item.color} />
                <Text style={styles.activityText}>{item.text}</Text>
                <Text style={styles.activityTime}>{item.time}</Text>
              </View>
              {i < 2 && <View style={styles.divider} />}
            </View>
          ))}
        </View>

        <View style={{ height: 120 }} />
      </ScrollView>

      {/* SOS Button */}
      <SOSButton />
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: "#0a0a0a" },
  scroll: { paddingHorizontal: 20, paddingTop: 60 },
  header: { flexDirection: "row", justifyContent: "space-between", alignItems: "flex-start", marginBottom: 24 },
  greeting: { color: "#666", fontSize: 14, fontWeight: "500" },
  appName: { color: "#fff", fontSize: 32, fontWeight: "900", letterSpacing: -1 },
  tagline: { color: "#444", fontSize: 12, marginTop: 2 },
  nightBtn: { width: 44, height: 44, borderRadius: 22, backgroundColor: "#111", justifyContent: "center", alignItems: "center", borderWidth: 1, borderColor: "#222" },
  nightBtnActive: { backgroundColor: "#60a5fa" },
  riskCard: { backgroundColor: "#111", borderRadius: 20, padding: 20, flexDirection: "row", justifyContent: "space-between", alignItems: "center", marginBottom: 12, borderWidth: 1, borderColor: "#fad047" + "44" },
  riskLeft: {},
  riskLabel: { color: "#f97316", fontSize: 10, fontWeight: "800", letterSpacing: 2 },
  riskValue: { color: "#fff", fontSize: 24, fontWeight: "800", marginTop: 4 },
  riskDesc: { color: "#666", fontSize: 12, marginTop: 2 },
  riskScore: { flexDirection: "row", alignItems: "flex-end" },
  riskNumber: { color: "#f97316", fontSize: 48, fontWeight: "900", lineHeight: 52 },
  riskMax: { color: "#444", fontSize: 16, marginBottom: 8 },
  tipCard: { backgroundColor: "#fad047" + "11", borderRadius: 14, padding: 14, flexDirection: "row", alignItems: "center", marginBottom: 24, borderWidth: 1, borderColor: "#fad047" + "33" },
  tipText: { color: "#fad047", fontSize: 13, marginLeft: 10, flex: 1, lineHeight: 18 },
  sectionTitle: { color: "#fff", fontSize: 18, fontWeight: "800", marginBottom: 14 },
  grid: { flexDirection: "row", flexWrap: "wrap", gap: 12, marginBottom: 24 },
  featureCard: { backgroundColor: "#111", borderRadius: 16, padding: 14, width: (width - 52) / 2, borderWidth: 1, borderColor: "#1a1a1a" },
  featureTop: { flexDirection: "row", justifyContent: "space-between", alignItems: "center", marginBottom: 10 },
  iconBox: { width: 40, height: 40, borderRadius: 12, justifyContent: "center", alignItems: "center" },
  badge: { paddingHorizontal: 6, paddingVertical: 2, borderRadius: 6 },
  badgeText: { fontSize: 9, fontWeight: "800", letterSpacing: 0.5 },
  featureTitle: { color: "#fff", fontSize: 14, fontWeight: "700", marginBottom: 4 },
  featureDesc: { color: "#555", fontSize: 11, lineHeight: 15 },
  nightBanner: { backgroundColor: "#60a5fa" + "11", borderRadius: 14, padding: 16, flexDirection: "row", alignItems: "center", marginBottom: 24, borderWidth: 1, borderColor: "#60a5fa" + "33" },
  nightBannerTitle: { color: "#60a5fa", fontWeight: "700", fontSize: 14 },
  nightBannerDesc: { color: "#60a5fa" + "99", fontSize: 12, marginTop: 2 },
  nightDot: { width: 8, height: 8, borderRadius: 4, backgroundColor: "#60a5fa" },
  activityCard: { backgroundColor: "#111", borderRadius: 16, padding: 16, marginBottom: 24, borderWidth: 1, borderColor: "#1a1a1a" },
  activityRow: { flexDirection: "row", alignItems: "center", paddingVertical: 4 },
  activityText: { color: "#ccc", fontSize: 13, flex: 1, marginLeft: 10 },
  activityTime: { color: "#444", fontSize: 11 },
  divider: { height: 1, backgroundColor: "#1a1a1a", marginVertical: 10 },
});